#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;
using Vec = vector<int>;

string findLongSubstring(string& str)
{
	int start = 0, len = 0;

	for (int i = 0; i < str.size(); i++)
	{
		int temp1 = i;
		int temp2 = i + 1;
		while (temp1 >= 0 and temp2 < str.size() and str[temp1] == str[temp2])
		{
			temp1--;
			temp2++;
		}

		if (len < ((temp2 - i) * 2))
		{
			start = temp1 + 1;
			len = (temp2 - 1 - i) * 2;
		}
		temp1 = i;
		temp2 = i + 2;
		while (temp1 >= 0 and temp2 < str.size() and str[temp1] == str[temp2])
		{
			temp1--;
			temp2++;
		}

		if (len < ((temp2 - 1 - i) * 2 + 1))
		{
			start = temp1 + 1;
			len = (temp2 - i) * 2 + 1;
		}

	}

	return str.substr(start, len);
}

int main()
{
	string a("ggabbahh");
	cout << findLongSubstring(a) << endl;
}


