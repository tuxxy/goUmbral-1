// Copyright (C) 2018 NuCypher
//
// This file is part of goUmbral.
//
// goUmbral is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// goUmbral is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with goUmbral. If not, see <https://www.gnu.org/licenses/>.
package crypto

import (
    "testing"
)

func TestError(t *testing.T) {
    bn1 := IntToBN(0)
    bn2 := IntToBN(5)

    bn3 := SubBN(bn1, bn2)

    bn4 := RandRangeBN(bn3)

    if bn4 != nil {
        t.Fail()
    }
}