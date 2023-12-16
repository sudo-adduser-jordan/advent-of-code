#include <iostream>
#include <vector>
#include <map>
#include <fstream>
#include <string>
#include "../lib/colors.cpp"
using namespace std;

bool isNumber(char const c)
{
    return c >= 48 && c <= 57;
}

bool isSymbol(char const c)
{
    return !isNumber(c) && c != '.';
}

void SumParts(int *sum, vector<string> const &vec, int const lines, int const lineLen)
{
    for (int row = 0; row < lines; row++)
    {
        int index = 0;
        int number = 0;
        bool isPart = false;

        while (index < lineLen)
        {
            char character = vec[row][index];

            if (!isNumber(character))
            {
                if (isPart)
                    *sum += number;
                number = 0;
                isPart = false;
            }

            if (isNumber(character))
            {
                number *= 10;
                number += character - '0'; // concat char to int

                if (index < lineLen - 1)
                    isPart = isPart || isSymbol(vec[row][index + 1]);

                if (index > 0)
                    isPart = isPart || isSymbol(vec[row][index - 1]);

                if (row < lines - 1)
                    isPart = isPart || isSymbol(vec[row + 1][index]);

                if (row < lines - 1 && index > 0)
                    isPart = isPart || isSymbol(vec[row + 1][index - 1]);

                if (row < lines - 1 && index < lineLen - 1)
                    isPart = isPart || isSymbol(vec[row + 1][index + 1]);

                if (row > 0)
                    isPart = isPart || isSymbol(vec[row - 1][index]);

                if (row > 0 && index > 0)
                    isPart = isPart || isSymbol(vec[row - 1][index - 1]);

                if (row > 0 && index < lineLen - 1)
                    isPart = isPart || isSymbol(vec[row - 1][index + 1]);
            }
            index++;
        }
        if (isPart)
            *sum += number;
    }
}

int main()
{
    vector<string> schematic = {};

    string str;
    ifstream file;
    file.open("input.txt");
    while (getline(file, str))
    {
        schematic.push_back(str);
    }
    file.close();

    int sum = 0;
    int lines = schematic.size();
    int lineLen = schematic[0].size();
    SumParts(&sum, schematic, lines, lineLen);

    cout << "The sum of all of the part numbers in the engine schematic is: " << GREEN_TEXT << sum << RESET_COLOR << endl;
}
