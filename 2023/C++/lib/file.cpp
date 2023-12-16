#include <string>
#include <fstream>
#include <sstream>
#include <vector>
using namespace std;

/**
 *  @brief  Read a text file into a string vector.
 *  @param  file Path to file.
 *  @return Copy of vector.
 *
 *  Creates an ifstream to read file line by line and push_back to fill vector.
 **/
vector<string> file_into_string_vector(const string &filename)
{
    vector<string> lines;

    string line;
    ifstream file(filename);
    while (getline(file, line))
    {
        lines.push_back(line);
    }

    return lines;
}