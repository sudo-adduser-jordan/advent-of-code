#include <string>
#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include "../lib/colors.cpp"
using namespace std;

// const string filename = "test.txt";
const string filename = "input.txt";

struct Card
{
    int cardNumber;
    vector<int> winningNumbers;
    vector<int> numbers;
    int totalMatches;
    int totalAmountOfCard;
};

int getPointTotal(Card &card)
{
    int totalMatches = 0;
    for (int winningNumber : card.winningNumbers)
        for (int number : card.numbers)
            if (number == winningNumber)
                totalMatches++;

    if (totalMatches == 0)
        return 0;

    if (totalMatches == 1)
        return 1;
    else
    {
        int total = 1;
        for (int i = 1; i < totalMatches; i++)
            total *= 2;

        return total;
    }
}

int getTotalMatches(Card &card)
{
    int totalMatches = 0;
    for (int winningNumber : card.winningNumbers)
        for (int number : card.numbers)
            if (number == winningNumber)
                totalMatches++;
    return totalMatches;
}

void PartOne(int *result, vector<Card> const &cards)
{
    for (Card card : cards)
        *result += getPointTotal(card);
}

void PartTwo(int *result, vector<Card> &cards)
{
    for (Card card : cards) // one of ea card w/ score
    {
        cards[card.cardNumber - 1].totalMatches = getTotalMatches(card);
        cards[card.cardNumber - 1].totalAmountOfCard = 1;
        *result += 1;
    }

    for (Card card : cards)
        while (card.totalAmountOfCard > 0)
        {
            for (int index = 0; index < card.totalMatches; index++)
            {
                cards[card.cardNumber + index].totalAmountOfCard++;
                *result += 1;
            }
            card.totalAmountOfCard--;
        }
}

int main()
{
    vector<Card> cards;

    string line;
    ifstream file(filename);
    while (getline(file, line))
    {
        string cardNumber = line.substr(line.find(' '), line.find(':') - line.find(' '));
        string winningNumbers = line.substr(line.find(':') + 1, line.find('|') - line.find(':') - 1);
        string numbers = line.substr(line.find('|') + 1);

        Card card;
        card.cardNumber = stoi(cardNumber);

        int number;
        stringstream streamTwo(winningNumbers);
        while (streamTwo >> number)
            card.winningNumbers.push_back(number);

        stringstream streamOne(numbers);
        while (streamOne >> number)
            card.numbers.push_back(number);

        cards.push_back(card);
    }

    int result = 0;
    // PartOne(&result, cards);
    // cout << "The total points of the pile of cards is: " << GREEN_TEXT << result << RESET_COLOR << endl;
    PartTwo(&result, cards);
    cout << "The amount of scratch cards won is: " << GREEN_TEXT << result << RESET_COLOR << endl;

    return 0;
}