# Check anagram in go

```sh
abc(3) abx(3) false map Took: 65ns
 abc(3) abx(3) false sort Took: 68ns
 abc(3) cab(3) true map Took: 131ns
 abc(3) cab(3) true sort Took: 70ns
 abc(3) caba(4) false map Took: 2ns
 abc(3) caba(4) false sort Took: 2ns
 abcdefghij(10) jihgfedcba(10) true sort Took: 182ns
 abcdefghij(10) jihgfedcba(10) true map Took: 796ns
 the quick brown fox(19) xof nworb kciuq eht(19) true map Took: 1.669µs
 the quick brown fox(19) xof nworb kciuq eht(19) true sort Took: 388ns
 aaaaaaaaaaaaaaaaaaaa(100) aaaaaaaaaaaaaaaaaaaa(100) true map Took: 1.911µs
 aaaaaaaaaaaaaaaaaaaa(100) aaaaaaaaaaaaaaaaaaaa(100) true sort Took: 1.604µs
 aaaaaaaaaaaaaaaaaaaa(1000) aaaaaaaaaaaaaaaaaaaa(1000) true map Took: 18.52µs
 aaaaaaaaaaaaaaaaaaaa(1000) aaaaaaaaaaaaaaaaaaaa(1000) true sort Took: 13.579µs
 aaaaaaaaaaaaaaaaaaaa(10000) aaaaaaaaaaaaaaaaaaaa(10000) true sort Took: 126.89µs
 aaaaaaaaaaaaaaaaaaaa(10000) aaaaaaaaaaaaaaaaaaaa(10000) true map Took: 184.036µs

```
