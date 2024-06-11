package bin

import (
	"reflect"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	tests := []struct {
		name        string
		bin         string
		data        BinData
		indexSearch []int
	}{
		{
			name: "Simple case",
			bin:  "0",
			data: BinData{
				isPrepaid: true,
			},
			indexSearch: []int{0},
		},
		{
			name: "complex case",
			bin:  "012345",
			data: BinData{
				isPrepaid: true,
			},
			indexSearch: []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		trie := NewBinRangeTrie()
		trie.Insert(test.bin, test.data)

		node := trie.root
		for _, indexSearch := range test.indexSearch {
			node = node.children[indexSearch]
		}
		if node.data != test.data {
			t.Errorf("wanted %v got %v", test.data, node.data)
		}
	}
}

func Test_Search(t *testing.T) {
	tests := []struct {
		name      string
		insertBin string
		searchBin string
		want      BinData
	}{
		{
			name:      "Simple case",
			insertBin: "0",
			searchBin: "0",
			want: BinData{
				isPrepaid: true,
			},
		},
		{
			name:      "Complex case",
			insertBin: "0123456",
			searchBin: "0123456",
			want: BinData{
				isPrepaid: true,
			},
		},
		{
			name:      "Missed result",
			insertBin: "0123456",
			searchBin: "012345",
			want: BinData{
				isPrepaid: true,
			},
		},
	}

	for _, test := range tests {
		trie := NewBinRangeTrie()
		trie.Insert(test.insertBin, test.want)

		got := trie.Search(test.searchBin)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("wanted %v got %v", test.want, got)
		}
	}
}
