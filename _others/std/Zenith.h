#ifndef Zenith_H
#define Zenith_H

#include <cstdint>
#include <cmath>
#include <array>
#include <vector>

namespace Zenith {
    using Bool = bool;

    using UInt8 = uint8_t;
    using UInt16 = uint16_t;
    using UInt32 = uint32_t;
    using UInt64 = uint64_t;
    using UInt = uint_fast32_t;

    using Int8 = int8_t;
    using Int16 = int16_t;
    using Int32 = int32_t;
    using Int64 = int64_t;
    using Int = int_fast32_t;

    using Float32 = float;
    using Float64 = double;

    template <typename T> using Slice = std::vector<T>;
    template <typename T, std::size_t N> using Array = std::array<T, N>;

    template <typename T>
    T pow(T left, T right) {
        if(right < 0 || left == 0) { return 0; }
        if(right == 0 || left == 1) { return 1; }

        T product = 1;

        while(right > 0) {
            if(right % 2 == 1) {
                product *= left;
                right--;
                continue;
            }

            left *= left;
            right /= 2;
        }

        return product;
    }
}

#endif // !Zenith_H
