# Get the latest tag and suggest an incremented version
LATEST_TAG=$(git tag --list "v*" | tail -n 1)

if [[ -n $LATEST_TAG ]]; then
    # Extract major, minor, and patch numbers
    IFS='.' read -r -a VERSION_PARTS <<< "${LATEST_TAG:1}"
    MAJOR="${VERSION_PARTS[0]}"
    MINOR="${VERSION_PARTS[1]}"
    PATCH="${VERSION_PARTS[2]}"

    # Increment the patch version
    PATCH=$((PATCH + 1))

    # Suggest the next tag
    SUGGESTED_TAG="v${MAJOR}.${MINOR}.${PATCH}"
else
    # If no tags exist, start with v1.0.0
    SUGGESTED_TAG="v1.0.0"
fi

echo "Suggested next tag: $SUGGESTED_TAG"
echo "Execute: git tag $SUGGESTED_TAG && git push $SUGGESTED_TAG"