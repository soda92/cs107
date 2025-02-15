#include <iostream>
#include <string>

using namespace std;

string s;

basic_string<int> a;

// sadjgwe
int GetInt()
{
  int x;
  cin >> x;
  return x;
}

enum directionT
{
  North,
  South,
  // fhwaegpwe
  East,
  West
};

directionT dir = East;

struct T0
{
  int a;
  char b;
};

#pragma pack(1)
struct T1
{
  int a;
  char b;
};
#pragma pop

double GetScoreAndAverage(int numScores);

int main()
{
  a += 11;
  cout << sizeof(T1) << ' ' << sizeof(T0) << '\n';
  cout << "Welcome" << endl;
  // double average = GetScoreAndAverage(4);
  // cout << "Average is " << average << endl;
  return 0;
}

double GetScoreAndAverage(int numScores)
{
  int sum;

  for (int i = 0; i < numScores; i++) {
    sum += GetInt();
  }
  return (double)sum / (double)numScores;
}