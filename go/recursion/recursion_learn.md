## 実際にコードを書いて再確認したこと

- `return`後のコードの評価
  - `return`後関数が実行されればまずは関数の結果が返ってくるまで評価を待機している
  - `return calc_sum(x) + cast_n`
- プログラムの挙動の前に具体的なケースを想定してみることが重要
  - フィボナッチ数列の具体ケースを書いてみるなど
  - フィボナッチ数列と何を最終的に出力するのかわかならい
- 処理の実際の流れを視覚化(図式，書き出す)ことをまずやる方が早くプログラミングできる

## フィボナッチ数列 1

### 概要

フィボナッチ数列は、最初の 2 つの数が「0」と「1」で始まります。その後の数は、前の 2 つの数を足して求められます。つまり、0、1、1（0 + 1）、2（1 + 1）、3（1 + 2）、5（2 + 3）、8（3 + 5）、13（5 + 8）

### フィボナッチ数列第 n 項を計算する具体例

- n が 3 の場合は 0, 1, 1, "2", 3 と続く数列の"2"を求めたい
  - 2: 1 + 0
  - 1: 1 + 0
  - 0: 0

### フィボナッチ数列の再帰

- トップダウンで n 項目を求めにいく
- f(n-1) + f(n-2)の各辺で最下層の 0 と 1 まで計算して各階層(各項目)の値を求める
- 合算していくと n 項目のフィボナッチ数列を表している
- 木構造で表した図を見るのが早い
  - [再帰アルゴリズムを使うとどう世界が広がるのか](https://qiita.com/drken/items/23a4f604fa3f505dd5ad)

### プログラムの挙動

- n が 3 の場合、以下のように計算が進みます
- calc_fibonacci(3) は calc_fibonacci(2) + calc_fibonacci(1) を計算します
- calc_fibonacci(2) は calc_fibonacci(1) + calc_fibonacci(0) を計算します
- calc_fibonacci(1) は 1 を返します
- calc_fibonacci(0) は 0 を返します
- 上記に処理により calc_fibonacci(2) は 1 + 0 = 1 を返します
- calc_fibonacci(3) は 1 + 1 = 2 を返します。

## 気づき

- 処理の実際の流れを視覚化(図式，書き出す)ことをまずやる方が早くプログラミングできる

## 参考

- [再帰アルゴリズムを使うとどう世界が広がるのか](https://qiita.com/drken/items/23a4f604fa3f505dd5ad)
- [末尾再帰による最適化](https://qiita.com/pebblip/items/cf8d3230969b2f6b3132)
