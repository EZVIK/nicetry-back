package main

/**
请你实现三个 API append，addAll和multAll来实现奇妙序列。

请实现Fancy类 ：

Fancy()初始化一个空序列对象。
void append(val) 将整数val添加在序列末尾。
void addAll(inc)将所有序列中的现有数值都增加inc。
void multAll(m)将序列中的所有现有数值都乘以整数m。
int getIndex(idx) 得到下标为idx处的数值（下标从 0 开始），并将结果对109 + 7取余。如果下标大于等于序列的长度，请返回-1。


输入：
["Fancy", "append", "addAll", "append", "multAll", "getIndex", "addAll", "append", "multAll", "getIndex", "getIndex", "getIndex"]
[[], [2], [3], [7], [2], [0], [3], [10], [2], [0], [1], [2]]
输出：
[null, null, null, null, null, 10, null, null, null, 26, 34, 20]

解释：
Fancy fancy = new Fancy();
fancy.append(2);   // 奇妙序列：[2]
fancy.addAll(3);   // 奇妙序列：[2+3] -> [5]
fancy.append(7);   // 奇妙序列：[5, 7]
fancy.multAll(2);  // 奇妙序列：[5*2, 7*2] -> [10, 14]
fancy.getIndex(0); // 返回 10
fancy.addAll(3);   // 奇妙序列：[10+3, 14+3] -> [13, 17]
fancy.append(10);  // 奇妙序列：[13, 17, 10]
fancy.multAll(2);  // 奇妙序列：[13*2, 17*2, 10*2] -> [26, 34, 20]
fancy.getIndex(0); // 返回 26
fancy.getIndex(1); // 返回 34
fancy.getIndex(2); // 返回 20

*/
