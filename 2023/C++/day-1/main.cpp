#include <stdlib.h>
#include <iostream>
#include <fstream>
#include <cmath>
#include <string>
#include <vector>
#include <cctype>
#include "../lib/colors.cpp"
using namespace std;

const string filename = "input.txt";
// const string filename = "test.txt";

void PartOne()
{
    int result;
    vector<int> calibrationValues;

    string text;
    ifstream file(filename);
    while (getline(file, text))
    {
        cout << text << endl;
        vector<char> digitVector;
        for (char character : text)
        {
            if (isdigit(character))
            {
                digitVector.push_back(character);
            }
        }

        char firstNumber = digitVector[0];
        char lastNumber = digitVector[digitVector.size() - 1];
        digitVector = {};

        string value;
        value += firstNumber;
        value += lastNumber;
        cout << value << "\n";

        calibrationValues.push_back(stoi(value));
    }
    file.close();

    for (int i : calibrationValues)
    {
        result += i;
    }

    cout << "\nThe sum of all of the calibration values is: " GREEN_TEXT << result << RESET_COLOR << "\n\n";
}

void PartTwo()
{
    int result;
    vector<int> calibrationValues;

    string text;
    ifstream file(filename);
    while (getline(file, text))
    {
        cout << text << endl;
        vector<char> digitVector;
        string s = "";
        for (char character : text)
        {

            if (isdigit(character))
            {
                digitVector.push_back(character);
                s = character;
            }
            else
            {
                s += character;
            }

            if (s.find("one") != std::string::npos)
            {
                digitVector.push_back('1');
                s = character;
            }
            if (s.find("two") != std::string::npos)
            {
                digitVector.push_back('2');
                s = character;
            }
            if (s.find("three") != std::string::npos)
            {
                digitVector.push_back('3');
                s = character;
            }
            if (s.find("four") != std::string::npos)
            {
                digitVector.push_back('4');
                s = character;
            }
            if (s.find("five") != std::string::npos)
            {
                digitVector.push_back('5');
                s = character;
            }
            if (s.find("six") != std::string::npos)
            {
                digitVector.push_back('6');
                s = character;
            }
            if (s.find("seven") != std::string::npos)
            {
                digitVector.push_back('7');
                s = character;
            }
            if (s.find("eight") != std::string::npos)
            {
                digitVector.push_back('8');
                s = character;
            }
            if (s.find("nine") != std::string::npos)
            {
                digitVector.push_back('9');
                s = character;
            }
        }

        char firstNumber = digitVector[0];
        char lastNumber = digitVector[digitVector.size() - 1];
        digitVector = {};

        string value;
        value += firstNumber;
        value += lastNumber;
        cout << value << "\n";

        calibrationValues.push_back(stoi(value));
    }
    file.close();

    for (int i : calibrationValues)
    {
        result += i;
    }

    cout << "\nThe sum of all of the calibration values is: " GREEN_TEXT << result << RESET_COLOR << "\n\n";
}

int main()
{
    // PartOne();
    PartTwo();
    return 0;
}