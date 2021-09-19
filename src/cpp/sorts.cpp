#include <fstream>
#include <iostream>
#include "time.h"

const int MAXSET = 10000;

// int (&set)[MAXSET] is passing the variable set by reference.
void initSet(int (&set)[MAXSET], std::string fileName="sampleset.txt") {
    std::ifstream inf;
    inf.open(fileName);
    
    int i = 0;
    while(!inf.eof() && i<MAXSET) {
        inf >> set[i];
        i++;
    }

    inf.close();
}

void printElements(int (&set)[MAXSET], std::ostream &out=std::cout) {
    out << "Index:\tElement:" << std::endl;
    for (int i=0; i<MAXSET/1000; i++)
        out << i*1000 << '\t' << set[i*1000] << std::endl;
    out << std::endl;
}

void printExecutionTime(clock_t &start, clock_t &end, std::ostream &out=std::cout) {
    // Calculates the execution time in milliseconds.
    long double time = (((long double)end-start)/CLOCKS_PER_SEC) * 100;
    out << "Execution time: " << time << "ms" << std::endl;
}

// The following is a bubble sort and quick sort algorithm implementation taken from a previous assignment.
void swap(int *x, int *y) {
    int tmp = *x;
    *x = *y;
    *y = tmp;
}

void bubbleSort(int set[MAXSET]) {
    for (int i=0; i<MAXSET-1; i++) {
        for (int j=0; j<MAXSET-i-1; j++) {
            if (set[j] > set[j+1]) {
                swap(&set[j], &set[j+1]);
            }
        }
    }
}

int partition(int (&set)[MAXSET], int low, int high) {
    int pivot = set[high];
    int i = (low - 1);
    for (int j=low; j<=high-1; j++) {
        if (set[j] < pivot) {
            i++;
            swap(&set[i], &set[j]);
        }
    }
    swap(&set[i+1], &set[high]);
    return (i + 1);
}

void quickSort(int (&set)[MAXSET], int start, int end) {
    if (start < end) {
        int pi = partition(set, start, end);

        quickSort(set, start, pi-1);
        quickSort(set, pi+1, end);
    }
}

void runSorts() {
    // C style arrays are faster to sort than the STL Array or Vector class, at least on my machine.
    // Also, Go's arrays are similar to C's arrays.
    // Since I am benchmarking here, I will stick with the data structure that gives me the best results.
    int set[MAXSET];
    initSet(set);
    clock_t start, end;
    
    std::cout << "[Bubble Sort]\n" << std::endl;
    std::cout << "Unsorted:" << std::endl;
    printElements(set);

    std::cout << "Sorted:" << std::endl;
    start = clock();
    bubbleSort(set);
    end = clock();
    printElements(set);
    printExecutionTime(start, end);

    // The set is changed directly in the functions to help speed up execution times.
    // Therefore the array is initialized again after each sort.
    initSet(set);
    std::cout << "\n[Quick Sort]\n" << std::endl;
    std::cout << "Unsorted:" << std::endl;
    printElements(set);

    std::cout << "Sorted:" << std::endl;
    start = clock();
    quickSort(set, 0, MAXSET-1);
    end = clock();
    printElements(set);
    printExecutionTime(start, end);
}

int main() {
    runSorts();
    return 0;
}
