# Buildozer example

    # This is executable Markdown that's tested on CI.
    set -o errexit -o nounset -o xtrace
    alias ~~~=":<<'~~~sh'";:<<'~~~sh'

## Try it out

~~~sh
output="$(aspect print //:all)"

# Verify that it produces the expected output
echo "${output}" | grep -q "all filegroup" || {
    echo >&2 "Wanted output containing 'all filegroup' but got '${output}'"
    exit 1
}
~~~
