#include <iostream>
using namespace std;

int main() {
    int num1, num2, result = 0;
    
    cout << "Enter two numbers: ";
    cin >> num1 >> num2;
    
    result = num1 + num2;
    
    cout << "The sum of " << num1 << " and " << num2 << " is " << result << endl;
    
    return 0;
}
