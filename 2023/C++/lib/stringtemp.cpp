#include <string>
#include <fstream>
#include <sstream>
#include <vector>
using namespace std;

/**
 *  @brief  Parse a string of ints into vector.
 *  @param  str String of ints.
 *  @return Copy of string.
 *
 *  Creates an isstream to read string of ints and pushback to fill vector.
 **/
vector<int> string_to_int_vector(const string &str)
{
    vector<int> ints;

    int number;
    stringstream iss(str);
    while (iss >> number)
        ints.push_back(number);

    return ints;
}

/**
 *  @brief  Find a string in another string.
 *  @param  match String to match.
 *  @param  str String to search.
 *  @return Boolean.
 *
 *  Uses string::find() and string::length() to search string and return a boolean value.
 **/
bool isFound(const string &match, const string &str)
{
    if (str.find(match) < str.length())
        return true;
    else
        return false;
}
