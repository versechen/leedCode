#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
using Vec = vector<int>;

int findKthNumber(vector<int> vec1, int vec1_start, vector<int> vec2, int vec2_start, int Kth)
{
	if (vec1_start + 1 > vec1.size()) return vec2[vec2_start + Kth - 1];
	if (vec2_start + 1 > vec2.size()) return vec1[vec1_start + Kth - 1];
	if (Kth == 1) return min(vec1[vec1_start], vec2[vec2_start]);
	int p1 = (vec1_start + Kth / 2) > vec1.size() ? INT32_MAX : vec1[vec1_start + Kth / 2 - 1];
	int p2 = (vec2_start + Kth / 2) > vec2.size() ? INT32_MAX : vec2[vec2_start + Kth / 2 - 1];
	if (p1 < p2)
		return findKthNumber(vec1, vec1_start + Kth / 2, vec2, vec2_start, Kth - Kth / 2);
	else
		return findKthNumber(vec2, vec2_start + Kth / 2, vec1, vec1_start, Kth - Kth / 2);
}

int main()
{
	vector<int> vec1{ 1,3,4,6,7,9 };
	vector<int> vec2{ 3,6,8,9,13,14 };
	int m = vec1.size();
	int n = vec2.size();
	double left = (findKthNumber(vec1, 0, vec2, 0, (m + n + 1) / 2));
	double right = (findKthNumber(vec1, 0, vec2, 0, (m + n + 2) / 2));
	cout << (left + right) / 2.0 << endl;
}


