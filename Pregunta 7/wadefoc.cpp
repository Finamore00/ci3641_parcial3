#include <boost/multiprecision/gmp.hpp>
#include <boost/math/special_functions/log1p.hpp>
#include <iostream>

using namespace boost::multiprecision;

mpz_int fib_aux(mpz_int x, mpz_int y, mpz_int k) {
    mpz_int p_res;
    return (k == 0 ? 0 : (k == 1 ? x : fib_aux(x+y, x, k-1)));
}

mpz_int fibonacci(mpz_int n) {
    return fib_aux(1, 0, n);
}

mpf_float my_log2(mpf_float x) {
    return boost::math::log1p(x-1)/boost::math::log1p(1);
}

mpz_int wadefoc(mpz_int n) {
    mpf_float nf = n.convert_to<boost::multiprecision::mpf_float>();
    mpf_float narayana = (((nf+1)*nf*nf*(n-1))/12);
    mpz_int l2 = my_log2(narayana + 1).convert_to<mpz_int>(); //Takes floor by converting
    return fibonacci(l2 + 1);
}

int main(int argC, char* argV[]) {
    mpz_int n = mpz_int(argV[1]);
    std::cout << wadefoc(n) << std::endl;
}