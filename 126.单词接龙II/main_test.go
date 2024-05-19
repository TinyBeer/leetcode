package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findLadders(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: `beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]`,
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: [][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"}},
		},
		{
			name: `beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]`,
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: nil,
		},
		// {
		// 	name: `beginWord = "qa", endWord = "sq", wordList =["si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"]`,
		// 	args: args{
		// 		beginWord: "qa",
		// 		endWord:   "sq",
		// 		wordList:  []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"},
		// 	},
		// 	want: [][]string{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList)
			sort.Slice(got, func(i, j int) bool {
				for x := 0; x < len(got[i]); x++ {
					if got[i][x] == got[j][x] {
						continue
					}
					return got[i][x] < got[j][x]
				}
				return false
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findLadders() = %v, want %v", got, tt.want)
			}
		})
	}
}
