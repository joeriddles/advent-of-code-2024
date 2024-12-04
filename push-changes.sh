if git diff --quiet; then
  exit 0
else
  git add .
  git commit -m "[Automated] Update README"
  git push
fi
