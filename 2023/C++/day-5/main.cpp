#include "../lib/colors.cpp" // Variables for colored output
#include <string>
#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <list>
#include <map>
#include <thread> // std::thread

#include <windows.h> // for windows

using namespace std;

// const string filename = "test.txt";
const string filename = "input.txt";

struct Range
{
    long long int destRangeStart;
    long long int sourceRangeStart;
    long long int RangeLen;
};

struct SeedRange
{
    long long int start;
    long long int end;
};

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

/**
 *  @brief  Parse a string of long long ints into vector.
 *  @param  str String of long long ints.
 *  @return Copy of string.
 *
 *  Creates an isstream to read string of ints and pushback to fill vector.
 **/
vector<long long int> string_to_long_long_vector(const string &str)
{
    vector<long long int> ints;

    long long int number;
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

void partOne(long long int *result, vector<long long int> &seeds, list<vector<Range>> const &listOfMaps)
{
    for (long long int &seed : seeds)        // for each seed
        for (vector<Range> map : listOfMaps) // for each map
            for (Range range : map)          // for each range in map
                if (range.sourceRangeStart <= seed &&
                    seed <= range.RangeLen + range.sourceRangeStart) // if seed in destination range
                {
                    seed = range.destRangeStart + (seed - range.sourceRangeStart); // convert seed to destination
                    break;
                }

    *result = seeds[0]; // return smallest value in array
    for (long long int seed : seeds)
        if (*result > seed)
            *result = seed;
}

void process_thread(SeedRange seedRange, list<vector<Range>> const listOfMaps, vector<long long int> &locations)
{

    vector<long long int> seeds;

    // fill seeds vector
    long long int index = seedRange.start;
    while (index < seedRange.end)
    {
        seeds.push_back(index);
        index++;
    }

    // process each seed
    for (long long int &seed : seeds)        // for each seed
        for (vector<Range> map : listOfMaps) // for each map
            for (Range range : map)          // for each range in map
                if (range.sourceRangeStart <= seed &&
                    seed <= range.RangeLen + range.sourceRangeStart) // if seed in destination range
                {
                    seed = range.destRangeStart + (seed - range.sourceRangeStart); // convert seed to destination
                    break;
                }

    long long int result = seeds[0];
    for (long long int seed : seeds)
        if (result > seed)
            result = seed;

    cout << "Thread finished for seed range: " << seedRange.start << endl;
    locations.push_back(result);
}

void partTwo(long long int *result, vector<long long int> &seedsVector, list<vector<Range>> const listOfMaps)
{
    vector<long long int> locations;

    long long int index = 0;
    vector<SeedRange> seedRanges;
    while (index < seedsVector.size())
    {
        SeedRange seedRange;
        seedRange.start = seedsVector[index];
        seedRange.end = seedsVector[index] + seedsVector[index + 1];

        seedRanges.push_back(seedRange);

        index++;
        index++;
    }

    // create threads
    vector<thread> threads;
    for (SeedRange seedrange : seedRanges)
    {
        thread td(process_thread, seedrange, listOfMaps, ref(locations));
        cout << "Thread created: " << td.get_id() << endl;
        threads.push_back(move(td));
    }

    // wait for threads to finish
    for (auto &t : threads)
    {
        // cout << "Thread finished: " << t.get_id() << endl;
        t.join();
    }

    // print locations
    *result = locations[0];
    for (long long int location : locations)
        if (*result > location)
            *result = location;
}

int main()
{
    // Initialize variables.
    long long int result = 0;                                                                          // End result.
    vector<string> input = file_into_string_vector(filename);                                          // File input.
    vector<long long int> seeds = string_to_long_long_vector(input[0].substr(input[0].find(':') + 1)); // List of seeds.
    list<vector<Range>> listOfMaps;                                                                    // List of Maps

    vector<Range> map;
    for (string line : input)
    {
        if (line.length() == 0)
            continue;

        if (isFound("map", line))
        {
            if (!map.empty())
                listOfMaps.push_back(map);
            map.clear();
            continue;
        }

        if (isdigit(line.at(0)))
        {
            vector<long long int> numbers = string_to_long_long_vector(line);

            Range range;
            range.destRangeStart = numbers[0];
            range.sourceRangeStart = numbers[1];
            range.RangeLen = numbers[2];

            map.push_back(range);
        }
    }
    listOfMaps.push_back(map);

    // Process Data.
    // partOne(&result, seeds, listOfMaps);
    partTwo(&result, seeds, listOfMaps);

    // Print result.
    std::cout << "The result is: " << GREEN_TEXT << result << RESET_COLOR << endl;
}
