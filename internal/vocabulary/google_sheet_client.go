package vocabulary

import (
	"context"
	"encoding/json"
	"os"
	"sort"

	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleSheetClient struct {
	Trie *Trie
}

func NewGoogleSheetClient(apiKey, spreadSheetID string, isVocabularyFileWillCreate bool) (*GoogleSheetClient, error) {
	svc, err := sheets.NewService(
		context.Background(),
		option.WithAPIKey(apiKey),
	)
	if err != nil {
		return nil, errors.Wrap(err, "error when google sheets new service")
	}

	resp, err := svc.Spreadsheets.
		Get(spreadSheetID).
		IncludeGridData(true).
		Do()
	if err != nil {
		return nil, errors.Wrap(err, "error when reading values from spreadsheet")
	}

	vocabularies := transformSheetResponse(resp)

	if isVocabularyFileWillCreate {
		return nil, createVocabularyFile(vocabularies)
	}

	tree := InitTrie()
	for _, voc := range vocabularies {
		tree.Insert(voc)
	}

	return &GoogleSheetClient{
		Trie: tree,
	}, nil
}

func (gsc *GoogleSheetClient) SuggestWordsByPrefix(ctx context.Context, prefix string) ([]*Vocabulary, error) {
	vocabularies := gsc.Trie.Suggest(ctx, prefix)
	return vocabularies, nil
}

func transformSheetResponse(resp *sheets.Spreadsheet) []Vocabulary {
	firstSheet := resp.Sheets[0]
	rows := firstSheet.Data[0].RowData

	vocabularies := make([]Vocabulary, 0, len(rows))
	wordToVocabulary := make(map[string]Vocabulary, len(rows))
	words := make([]string, 0, len(rows))

	for _, row := range rows {
		cellData := row.Values

		voc := Vocabulary{}
		if len(cellData) > 0 {
			voc.Word = cellData[0].FormattedValue
		}
		if len(cellData) > 1 {
			voc.Meaning = cellData[1].FormattedValue
		}
		if len(cellData) > 2 {
			voc.Sentence = cellData[2].FormattedValue
		}

		wordToVocabulary[voc.Word] = voc
		words = append(words, voc.Word)
	}

	sort.Strings(words)

	for i := range words {
		vocabularies = append(vocabularies, wordToVocabulary[words[i]])
	}

	return vocabularies
}

func createVocabularyFile(vocabularies []Vocabulary) error {
	vocabularyFile, err := os.Create("internal/vocabulary/vocabulary.json")
	if err != nil {
		return errors.Wrap(err, "error when creating vocabulary json")
	}
	defer vocabularyFile.Close()

	vocabularyBytes, err := json.Marshal(vocabularies)
	if err != nil {
		return errors.Wrap(err, "error when marshaling vocabularies")
	}
	_, _ = vocabularyFile.Write(vocabularyBytes)

	return nil
}
