#include <iostream>
#include <fstream>
#include <string>

int part1() {
    int sum = 0, max = 0; 
    std::string line; 
    std::ifstream input("input.txt");
    while (!input.eof()) {
        line = "";
        std::getline(input, line);
        if (line != "") 
            sum += std::stoi(line);
        else {
            if (sum > max) 
                max = sum;
            sum = 0;
        }
    }
    return max;
}

int part2(int* arr, int n) {
    if (n == 3)
        return arr[0] + arr[1] + arr[2];
    int sum = 0, max = 0; 
    std::string line; 
    std::ifstream input("input.txt");
    while (!input.eof()) {
        line = "";
        std::getline(input, line);
        if (line != "") 
            sum += std::stoi(line);
        else {
            if (sum > arr[n] && sum != arr[0] && 
                sum != arr[1] && sum != arr[2]) 
                arr[n] = sum;
            sum = 0;
        }
    }
    return part2(arr, n+1);
}

int main(void) {
    int array[] = {0, 0, 0};
    std::cout << "Solution to part one => " << part1() << std::endl;
    std::cout << "Solution to part two => " << part2(array, 0) << std::endl;
}