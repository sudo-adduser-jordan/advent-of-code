#include <stdlib.h>
#include <string>
#include <iostream>
#include <fstream>
#include <cmath>
#include <vector>
#include <cctype>
#include <map>
#include <algorithm>
#include "../lib/colors.cpp"
using namespace std;

struct Set
{
    int red;
    int green;
    int blue;
};

struct Game
{
    int game;
    vector<Set> sets;
};

struct Results
{
    vector<Game> results;
};

// const string filename = "test.txt";
const string filename = "input.txt";
const Set loadout = {12, 13, 14};

void PartOne()
{
    Results results;

    string text;
    ifstream file(filename);
    while (getline(file, text))
    {
        Set set;
        set.red = 0;
        set.blue = 0;
        set.green = 0;

        Game game;
        game.game = stoi(text.substr(text.find(" "), text.find(":")));

        string str = text.substr(text.find(": ") + 2, text.length()) + ";";
        while (str.find("red") != string::npos)
        {
            str.replace(str.find("red"), sizeof("red") - 1, "r");
        }
        while (str.find("green") != string::npos)
        {
            str.replace(str.find("green"), sizeof("green") - 1, "g");
        }
        while (str.find("blue") != string::npos)
        {
            str.replace(str.find("blue"), sizeof("blue") - 1, "b");
        }

        int num;
        string value;
        for (char character : str)
        {
            if (isdigit(character))
            {
                value += character;
                num = stoi(value);
            }

            switch (character)
            {
            case ';':
                game.sets.push_back(set);
                set.red = 0;
                set.blue = 0;
                set.green = 0;
                value = "";
                num = 0;
                break;
            case 'r':
                set.red = num;
                value = "";
                num = 0;
                break;
            case 'g':
                set.green = num;
                value = "";
                num = 0;
                break;
            case 'b':
                set.blue = num;
                value = "";
                num = 0;
                break;
            default:
                break;
            }
        }
        results.results.push_back(game);
    }
    file.close();

    vector<int> possibleGames;
    for (Game game : results.results)
    {
        bool isValid = true;
        for (Set set : game.sets)
        {
            if (set.red > loadout.red || set.green > loadout.green || set.blue > loadout.blue)
            {
                isValid = false;
            }
        }

        if (isValid)
        {
            possibleGames.push_back(game.game);
        }
    }

    int sum = 0;
    for (int num : possibleGames)
    {
        cout << num << endl;
        sum += num;
    }

    cout << "The sum of the IDs of the possible games is: " << GREEN_TEXT << sum << RESET_COLOR << endl;
}

void PartTwo()
{
    Results results;

    string text;
    ifstream file(filename);
    while (getline(file, text))
    {
        Set set;
        set.red = 0;
        set.blue = 0;
        set.green = 0;

        Game game;
        game.game = stoi(text.substr(text.find(" "), text.find(":")));

        string str = text.substr(text.find(": ") + 2, text.length()) + ";";
        while (str.find("red") != string::npos)
        {
            str.replace(str.find("red"), sizeof("red") - 1, "r");
        }
        while (str.find("green") != string::npos)
        {
            str.replace(str.find("green"), sizeof("green") - 1, "g");
        }
        while (str.find("blue") != string::npos)
        {
            str.replace(str.find("blue"), sizeof("blue") - 1, "b");
        }

        int num;
        string value;
        for (char character : str)
        {
            if (isdigit(character))
            {
                value += character;
                num = stoi(value);
            }

            switch (character)
            {
            case ';':
                game.sets.push_back(set);
                set.red = 0;
                set.blue = 0;
                set.green = 0;
                value = "";
                num = 0;
                break;
            case 'r':
                set.red = num;
                value = "";
                num = 0;
                break;
            case 'g':
                set.green = num;
                value = "";
                num = 0;
                break;
            case 'b':
                set.blue = num;
                value = "";
                num = 0;
                break;
            default:
                break;
            }
        }
        results.results.push_back(game);
    }
    file.close();

    vector<int> powers;
    for (Game game : results.results)
    {
        int red = 0;
        int green = 0;
        int blue = 0;

        for (Set set : game.sets)
        {
            if (red < set.red)
            {
                red = set.red;
            }
            if (green < set.green)
            {
                green = set.green;
            }
            if (blue < set.blue)
            {
                blue = set.blue;
            }
        }

        int power = red * green * blue;
        powers.push_back(power);
    }

    int sum = 0;
    for (int n : powers)
    {
        sum += n;
    }

    cout << "The sum of the power of the sets is: " << GREEN_TEXT << sum << RESET_COLOR << endl;
}

int main()
{
    // PartOne();
    PartTwo();
    return 0;
}