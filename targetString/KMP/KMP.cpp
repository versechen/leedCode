#include <iostream>
#include <vector>
using namespace std;

void preKmp(const char* p, vector<int>& next)//fill next，目标与模式某位不匹配时，利用已有的最大前缀进行移动，减少对比
{
    int i = 0;// 一直增大，负责填满next
    int j = next[0] = -1;//负责模式，代表最大前缀的后一位
    while (i < next.size()-1)//保证next不溢出，保证终止条件
    {
        while (j > -1 && p[i] != p[j])//找到i对应的最大前缀
        {
            j = next[j];//确定最大前缀位
        }
        i++;//在新的起始位前进
        j++;//在新的起始位前进

        /*填写next[i]*/
        if (p[i] == p[j])
        {
            next[i] = next[j];//如果相等，带来的问题是模式新的对比位仍然和上次的字母一样，说明必定和目标不一样，需要继续变位
        }
        else
        {
            next[i] = j;//确定next[i]的数值，*退出时，i为size大小，保留的是整个模式的最大有效前缀。
        }
    }
}

void KMP(string p, string t)
{
    size_t m = p.size();
    size_t n = t.size();
    vector<int> kmpNext(m+1,0);//注意，size会大1,目的是最后一位存整个目标的最大前缀
    preKmp(p.c_str(),kmpNext);
    
    int i = 0;
    int j = 0;

    while (i < n)
    {
        while (j > -1 && p[j] != t[i])
        {
            j = kmpNext[j];//改变最大前缀的位置，保证新的p[j]不等于旧的p[j],但是未必和t[i]相等
        }

        j++;//i之前的已经相等，前进
        i++;//j之前的已经相等，前进

        if (j >= m)//如果j到了有效位的后一位，说明已经完全配置上了
        {
            cout << "find same string in t, index is " << i - j << endl;//输出目标匹配的起始索引
            j = kmpNext[j];//j是整个模式的最大前缀。继续进行对比，会复用已经匹配上的
        }
    }
}

int main()
{
//    string a = "hihihiuuuoj";
//    string b = "uu";
    string a;
    string b;
    cin >> a;
    cin >> b;

    KMP(b, a);
    std::cout << "end" << std::endl;

}

