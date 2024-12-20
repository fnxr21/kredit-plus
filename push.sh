git add .

echo 'Enter the commit message:'

read commitMessage

git commit -m "$commitMessage"

echo 'Enter the branch name:'

git branch -v

read branch

git remote -v

echo 'Enter the commit message:'

read remote

git push "${remote}" "${branch}"

echo 'Thank you for updated your repository'
