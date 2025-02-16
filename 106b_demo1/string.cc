#include <iostream>
#include <string>

using namespace std;

int main()
{
  string s = "abcde";

  auto s2 = s.substr(0, 3);

  cout << s2 << endl;
}