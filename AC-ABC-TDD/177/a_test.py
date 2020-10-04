from a import is_intime


def test_is_intime(self):
    cases = [
        [1000, 15, 80, True],
        [2000, 20, 100, True],
        [1000, 1, 1, False]
    ]
    for d, t, s, expected in cases:
        assert is_intime(d, t, s) == expected
