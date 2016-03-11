function solution(S) {

    if (S.length === 0) {
        return 1;
    }

    var otags = ['{', '[', '('];
    var ctags = ['}', ']', ')'];

    var list = S.split('');
    var nest = [];

    for (var i = 0; i < S.length; i++) {

        var oi = otags.indexOf(S[i]);
        var ci = ctags.indexOf(S[i]);

        if (oi != -1) {
            if (i == S.length - 1) {
                return 0;
            }
            nest.push(S[i]);
        } else if (ci != -1) {

            if (nest.length === 0) {
                return 0;
            }

            if (nest[nest.length - 1] != otags[ci]) {
                return 0;
            }
            nest.pop();
        }
    }

    return nest.length === 0 ? 1 : 0;
}