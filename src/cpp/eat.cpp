#include <iostream>
#include <string>
#include <vector>
#include <cmath>
#include "time.h"

using namespace std;

void printResults(
    const int &a, 
    const int &e, 
    const int &h, 
    const int &l, 
    const int &p, 
    const int &t,
    const int &eat,
    const int &that,
    const int &apple
) {
    cout << "A: " << a << ' ';
    cout << "E: " << e << ' ';
    cout << "H: " << h << ' ';
    cout << "L: " << l << ' ';
    cout << "P: " << p << ' ';
    cout << "T: " << t << endl;

    cout << "EAT: " << eat << ' ';
    cout << "THAT: " << that << ' ';
    cout << "APPLE: " << apple << endl;
}

// Helper function to print execution time.
void printExecutionTime(clock_t &start, clock_t &end, std::ostream &out=std::cout) {
    // Calculates the execution time in milliseconds.
    long double time = (((long double)end-start)/CLOCKS_PER_SEC) * 100;
    out << "Execution time: " << time << "ms" << std::endl;
}

void digitizeEat(const int &eat, int &e, int &a, int &t) {
    int rem = eat;
    e = rem / 100;
    rem %= 100;
    a = rem / 10;
    t = rem % 10;
}

void digitizeThat(const int &that, int &t1, int &h, int &a, int &t2) {
    int rem = that;
    t1 = rem / 1000;
    rem %= 1000;
    h = rem / 100;
    rem %= 100;
    a = rem / 10;
    t2 = rem % 10;
}

void digitizeApple(const int &apple, int &a, int &p1, int &p2, int &l, int &e) {
    int rem = apple;
    a = rem / 10000;
    rem %= 10000;
    p1 = rem / 1000;
    rem %= 1000;
    p2 = rem / 100;
    rem %= 100;
    l = rem / 10;
    e = rem % 10;
}

void findApple() {
    bool foundApple = false;
    int eat, e1, a1, t1;
    int that, t21, h2, a2, t22;
    int apple, a3, p31, p32, l3, e3;

    for (eat=100; eat<1000; eat++) {
        digitizeEat(eat, e1, a1, t1);
        if (e1 != a1 && e1 != t1 && a1 != t1) {
            for (that=1000; that<10000; that++) {
                digitizeThat(that, t21, h2, a2, t22);
                if (t22 == t21 && t1 == t21 && a1 == a2 && t21 != h2 && h2 != a2) {
                    if (eat + that >= 10000) {
                        apple = eat + that;
                        digitizeApple(apple, a3, p31, p32, l3, e3);
                        if (a1 == a3 && p31 == p32 && e1 == e3 && t21 != l3 && a3 != p31 && a3 != l3 && a3 != e3 && p31 != l3 && p31 != e3 && l3 != e3) {
                            printResults(a1, e1, h2, l3, p31, t1, eat, that, apple);
                            foundApple = true;
                        } else continue;
                    } else continue;
                } else continue;

                if (foundApple) break;
            }
        } else continue;

        if (foundApple) break;
    }
}

int main() {
    clock_t start, end;

    start = clock();
    findApple();
    end = clock();
    printExecutionTime(start, end);

    return 0;
}
