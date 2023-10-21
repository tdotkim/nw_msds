@echo OFF
cd "runs"
del *.*

cd ".."

FOR /L %%y IN (1, 1, 100) DO (
    go run main.go > NUL
    echo "go run number %%y"
)

FOR /L %%y IN (1, 1, 100) DO (
    python miller_py.py > NUL
    echo "py run number %%y"
)

FOR /L %%y IN (1, 1, 100) DO (
    "C:\Program Files\R\R-4.3.1\bin\Rscript.exe" miller_r.R > NUL
    echo "r run number %%y"
)

python results.py

PAUSE