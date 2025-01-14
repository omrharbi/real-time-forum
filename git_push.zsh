#!/bin/zsh

# Check if there are any staged or untracked changes
if [[ -z $(git status --porcelain) ]]; then
    echo "No changes to commit."
    exit 0
fi

# Prompt for commit message
echo -n "Enter commit message: "
read commit_message

# Add all changes
echo "Adding changes..."
git add .

# Commit changes
echo "Committing changes..."
git commit -m "$commit_message"

# Push to the remote repository
echo "Pushing to GitHub..."
git push origin main

echo "Changes pushed successfully."
