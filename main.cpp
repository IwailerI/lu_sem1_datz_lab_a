/*
Andrejs Piroženoks, ap24069
A26. Given 3 positive numbers. Check, does exist triangle with sides of given
length. (C++)
Program created: 2024/09/12

Notes: full documentation is provided using doxygen.
*/

#include <iostream>
#include <string>

/// @brief Reports weather a and b are close enough together, to be considered
/// equal.
bool float_eq(double a, double b) {
    // epsilon value is chosen arbitrarily
    const double EPS = 1e-8;
    return std::abs(a - b) < EPS;
}

/// @brief Reports weather triangle with given side lengths a, b, c can exist.
/// @param a side length a
/// @param b side length b
/// @param c side length c
/// @return true if triangle can exist, false otherwise.
bool triangle_exists(double a, double b, double c) {
    // this check is not strictly required, because user input is validated
    // but it is still here, just in case
    if (a <= 0.0 || b <= 0.0 || c <= 0.0) {
        return false;
    }

    // if sum of 2 sides is approximately the third, triangle cannot exist
    // this is here to mitigate floating-point error
    if (float_eq(a + b, c) || float_eq(b + c, a) || float_eq(a + c, b)) {
        return false;
    }

    // triangle exists if sum of any two sides' lengths is greater then the
    // third side's length
    return (a + b) > c && (a + c) > b && (b + c) > a;
}

/// @brief Prompts user for a boolean value.
///
/// Multiple styles of answers are supported: 0/1, y/n, t/f.
/// User will be re-prompted on invalid input.
/// No prompt message is provided on the first attempt.
/// @return User provided boolean value.
bool get_input_bool() {
    const char* ERROR_MESSAGE =
        "Invalid format. Accepted values for YES are 'Y', '1' or "
        "'T'. Accepted values for NO are 'N', '0' or 'F'. Please "
        "try again.";

    while (true) {
        std::string input_string;

        std::cin >> input_string;

        if (input_string.length() != 1) {
            std::cout << ERROR_MESSAGE << std::endl;
            continue;
        }

        char input = input_string[0];
        input = tolower(input);

        switch (input) {
            case '0':
            case 'f':
            case 'n':
                return false;

            case '1':
            case 't':
            case 'y':
                return true;

            default:
                break;
        }

        std::cout << ERROR_MESSAGE << std::endl;
    }
}

/// @brief Prompts user for a positive real value.
/// @note Does not handle invalid input.
/// @return User provided input.
double get_input_double() {
    while (true) {
        double input;

        std::cin >> input;

        if (input <= 0) {
            std::cout << "Length must be positive, please try again."
                      << std::endl;
            continue;
        }

        return input;
    }
}

/// @brief Wrapper for triangle_exists, that prompts user for input and
/// outputs the result.
void calculation() {
    std::cout << "Please enter side length a." << std::endl;
    double a = get_input_double();

    std::cout << "Please enter side length b." << std::endl;
    double b = get_input_double();

    std::cout << "Please enter side length c." << std::endl;
    double c = get_input_double();

    if (triangle_exists(a, b, c)) {
        std::cout << "Triangle exists." << std::endl;
    } else {
        std::cout << "Triangle doesn't exist." << std::endl;
    }
}

#ifdef SIMPLE_IO

// This version of main only contains a bare-bones I/O for automated testing.
int main() {
    double a, b, c;
    std::cin >> a >> b >> c;
    std::cout << triangle_exists(a, b, c) << std::endl;
}

#else

// This version of main contains full user I/O along with error handling.
int main() {
    bool do_another;
    do {
        calculation();
        std::cout << "Continue? (Y/N)" << std::endl;
        do_another = get_input_bool();
    } while (do_another);
}

#endif