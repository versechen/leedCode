#include <iostream>

const char* strstr_mine(const char* src,const char* dest)
{
    if (!src||!dest)
        return NULL;

    const char* src_rep = src;
    const char* dest_rep = dest;

    while (*src) //src 不断增加，末尾为\0，解析为false
    {
        src_rep = src;//保存新的源字符串的起始地址，以便返回
        dest_rep = dest;//初始化目标字符串起始位置
        do
        {
            if (!*dest_rep)//目标字符串完全匹配后，会到最后的\0,认为已经搜索到了结果；
                return src;
        } while (*src_rep++ == *dest_rep++);//如果对比的字符相等，继续往后移位
        src++;//没有返回，则说明src的其起始位置没有找到，需要将src往后移位
    }

    return NULL;
}

int main()
{
    const char* src = "hello";
    const char* dest = "ll";
    std::cout << strstr_mine(src, dest) - src << std::endl;
};

/*
原理：
在src中不断改变初始对比位置，如果比到目标字符的最后\0位，说明src存在相同子串，否则src++，重新对比
*/



