#include <fstream>
#include <iostream>

using namespace std;

void Error(string s)
{
  cerr << s << '\n';
  exit(-1);
}

int main()
{
  ifstream in;
  cout << "enter name: ";
  string s;
  string::reverse_iterator rev = s.rbegin();
  std::getline(cin, s);
  in.open(s);

  if (in.fail())
    Error("file didn't open");
}