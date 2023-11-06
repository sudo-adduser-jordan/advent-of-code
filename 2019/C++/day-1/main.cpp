#include <iostream>
#include <fstream>
#include <stdlib.h>
#include <cmath>
#include <string>
#include "../lib/colors.cpp"
using namespace std;

string filename = "input.txt";

int FindFuelCapacity(int mass)
{
    return round(mass / 3) - 2;
}

int FindFuelRequired(int mass)
{
    int total_fuel = round(mass / 3) - 2;
    int fuel = round(mass / 3) - 2;

    while (fuel >= 3)
    {
        fuel = round(fuel / 3) - 2;
        if (fuel > 0)
        {
            total_fuel = total_fuel + fuel;
            cout << "mass: " << mass << "\n";
            cout << "fuel: " << fuel << "\n";
            cout << "total fuel: " << total_fuel << "\n";
        }
    }
    return total_fuel;
}

void PrintResult(int result)
{
    cout << "\nTotal units of fuel required: " GREEN_TEXT << result << RESET_COLOR << "\n\n";
}

void PartOne()
{
    string text;
    ifstream file(filename);
    int result = 0;
    while (getline(file, text))
    {
        cout << MAGENTA_TEXT << text << RESET_COLOR << "\n";
        int mass = stoi(text);
        int fuel_capacity = FindFuelCapacity(mass);
        result = result + fuel_capacity;
    }
    file.close();
    PrintResult(result);
}

void PartTwo()
{
    string text;
    ifstream file(filename);
    int result = 0;
    while (getline(file, text))
    {
        cout << MAGENTA_TEXT << text << RESET_COLOR << "\n";
        int mass = stoi(text);
        int fuel_req = FindFuelRequired(mass);
        result = result + fuel_req;
        cout << GREEN_TEXT << result << RESET_COLOR << "\n";
    }
    file.close();
    PrintResult(result);
}

int main()
{
    PartOne();
    PartTwo();
    return 0;
}