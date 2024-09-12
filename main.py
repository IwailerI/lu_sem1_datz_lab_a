# Andrejs Piroženoks, ap24069
# A26. Given 3 positive numbers. Check, does exist triangle with sides of given
# length. (Python)
# Program created: 2024/09/12

# Notes: full documentation is provided using reStructuredText.


# I swear I am not doing anything malicious, I only need command line arguments.
from sys import argv


def float_eq(a: float, b: float) -> bool:
    EPS = 1e-8
    return abs(a - b) < EPS


def triangle_exists(a: float, b: float, c: float) -> bool:
    """Reports whether triangle with side lengths a, b, c can exist.

    :param a: length of side a
    :type a: float
    :param b: length of side b
    :type b: float
    :param c: length of side c
    :type c: float
    :return: true if triangle can exist, false otherwise
    :rtype: bool
    """

    # this check is not strictly required, because user input is validated
    # but it is still here, just in case
    if a <= 0 or b <= 0 or c <= 0:
        return False

    # if sum of 2 sides is approximately the third, triangle cannot exist
    # this is here to mitigate floating-point error
    if float_eq(a + b, c) or float_eq(a + c, b) or float_eq(b + c, a):
        return False

    # triangle exists if sum of any two sides' lengths is greater then the
    # third side's length
    return (a + b) > c and (a + c) > b and (b + c) > a


def input_bool() -> bool:
    """Prompts user for a boolean value.
    Accepts multiple value styles: 1/0, t/f, y/n.
    Prints a message on errors and prompts user again.
    Does not print any message on first prompt.

    :return: boolean value provided by the user.
    :rtype: bool
    """

    ERROR_MESSAGE = (
        "Invalid format. Accepted values for YES are 'Y', '1' or 'T'. "
        "Accepted values for NO are 'N', '0' or 'F'. Please try again."
    )

    while True:
        text = input().strip().lower()
        if text in "y1t":
            return True
        elif text in "n0f":
            return False
        print(ERROR_MESSAGE)


def input_float() -> float:
    """Prompts user for a positive real value.
    Prints a message on errors and prompts user again.
    Does not print any message on first prompt.

    :return: positive real value provided by the user.
    :rtype: float
    """
    ERROR_MESSAGE = "Invalid value. Please input a positive real number."

    while True:
        text = input().strip().lower()
        try:
            parsed = float(text)
        except ValueError:
            print(ERROR_MESSAGE)
            continue

        if parsed <= 0.0:
            print(ERROR_MESSAGE)
            continue

        return parsed


def calculation() -> None:
    """Wrapper function around triangle_exists, that asks user for parameters
    and prints the result.
    """
    print("Please enter side length a.")
    a = input_float()
    print("Please enter side length b.")
    b = input_float()
    print("Please enter side length c.")
    c = input_float()

    if triangle_exists(a, b, c):
        print("Triangle exists.")
    else:
        print("Triangle doesn't exist.")


def main() -> None:
    while True:
        calculation()
        print("Continue? (Y/N)")
        if not input_bool():
            break


def stripped_main() -> None:
    inp = input().split(maxsplit=2)
    data = [float(inp[i]) for i in range(3)]
    result = triangle_exists(*data)
    print(int(result))


if __name__ == "__main__":
    if len(argv) >= 2 and "--stripped" in argv[1:]:
        stripped_main()
    else:
        main()
