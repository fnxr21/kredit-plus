git add .

echo 'Enter the commit message:'

read commitMessage

git commit -m "$commitMessage"

echo 'Enter the branch name:'

git branch -v

read branch

git remote -v

echo 'Enter the commit message:'


git push origin dev

echo 'Thank you for updated your repository'
