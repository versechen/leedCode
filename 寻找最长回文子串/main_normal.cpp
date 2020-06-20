#include<iostream>
#include<vector>
#include<algorithm>
#include<thread>
using namespace std;
using Vec = vector<int>;

int main()
{
	const int limits = thread::hardware_concurrency();
	cout << limits << endl;
}


