function partialDot(val1, val2) {
    return val1*val2;
}

function dotProduct(v1, v2) {

    local threads = [];
    local accum = 0;
    local partialRes = [];

    if (v1.len() != v2.len()) {
        ::print("Los vectores no son del mismo tamaño");
        return null;
    }

    //Create thread pool
    for (local i = 0; i < v1.len(); i++) {
        threads.append(::newthread(partialDot));
    }

    for (local i = 0; i < v1.len(); i++) {
        accum += threads[i].call(v1[i], v2[i]);
    }

    return accum;

}

local v1 = [1, 2, 3, 4];
local v2 = [6, 7, 8, 9];
::print(::dotProduct(v1, v2) + "\n");