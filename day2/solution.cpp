#include <iostream>
#include <fstream>
#include <string>

#define ROCK_L     'A'
#define PAPER_L    'B'
#define SCISSORS_L 'C'

#define ROCK_R     'X'
#define PAPER_R    'Y'
#define SCISSORS_R 'Z'

#define ROCK_SCORE     1
#define PAPER_SCORE    2
#define SCISSORS_SCORE 3

#define LOSS_SCORE 0
#define DRAW_SCORE 3
#define WIN_SCORE  6

int part1() {
    std::ifstream input("input.txt");
    std::string line;
    char left, right;
    int score = 0;
    while(!input.eof()) {
        line = "";
        std::getline(input, line);
        left = line[0], right = line[2];
        switch (left) {
            case ROCK_L:
                switch (right) {
                    case ROCK_R:
                        score += ROCK_SCORE + DRAW_SCORE;
                        break;
                    case PAPER_R:
                        score += PAPER_SCORE + WIN_SCORE;
                        break;
                    case SCISSORS_R:
                        score += SCISSORS_SCORE + LOSS_SCORE;
                        break;
                }
                break;
            case PAPER_L:
                switch (right) {
                    case ROCK_R:
                        score += ROCK_SCORE + LOSS_SCORE;
                        break;
                    case PAPER_R:
                        score += PAPER_SCORE + DRAW_SCORE;
                        break;
                    case SCISSORS_R:
                        score += SCISSORS_SCORE + WIN_SCORE;
                        break;
                }
                break;
            case SCISSORS_L:
                switch (right) {
                    case ROCK_R:
                        score += ROCK_SCORE + WIN_SCORE;
                        break;
                    case PAPER_R:
                        score += PAPER_SCORE + LOSS_SCORE;
                        break;
                    case SCISSORS_R:
                        score += SCISSORS_SCORE + DRAW_SCORE;
                        break;
                }
                break;
        }
    }
    return score;
}

#define LOSE 'X'
#define DRAW 'Y'
#define WIN  'Z'

int part2() {
    std::ifstream input("input.txt");
    std::string line;
    char left, right;
    int score = 0;
    while(!input.eof()) {
        line = "";
        std::getline(input, line);
        left = line[0], right = line[2];
        switch (right) {
            case LOSE:
                switch (left) {
                    case ROCK_L:
                        score += SCISSORS_SCORE;
                        break;
                    case PAPER_L:
                        score += ROCK_SCORE;
                        break;
                    case SCISSORS_L:
                        score += PAPER_SCORE;
                        break;
                }
                score += LOSS_SCORE;
                break;
            case DRAW:
                switch (left) {
                    case ROCK_L:
                        score += ROCK_SCORE;
                        break;
                    case PAPER_L:
                        score += PAPER_SCORE;
                        break;
                    case SCISSORS_L:
                        score += SCISSORS_SCORE;
                        break;
                }
                score += DRAW_SCORE;
                break;
            case WIN:
                switch (left) {
                    case ROCK_L:
                        score += PAPER_SCORE;
                        break;
                    case PAPER_L:
                        score += SCISSORS_SCORE;
                        break;
                    case SCISSORS_L:
                        score += ROCK_SCORE;
                        break;
                }
                score += WIN_SCORE;
                break;
        }
    }
    return score;
}

int main() {
    std::cout << "Solution to part one => " << part1() << std::endl;
    std::cout << "Solution to part two => " << part2() << std::endl;
}