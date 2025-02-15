#include <iostream>
#include <string>

using namespace std;

void RemoveOccurrences(char ch, string& s)
{
  int found;
  while ((found = s.find(ch)) != string::npos) {
    s.erase(found, 1);
  }
}

int main()
{
  string myString = "chihuahua cheese crackers";
  RemoveOccurrences('c', myString);
  cout << myString << endl;
  return 0;
}