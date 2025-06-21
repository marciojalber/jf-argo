@echo off
cls

:: Marca o início do script
set h1=%time:~0,2%
if "%h1:~0,1%"==" " set h1=0%h1:~1,1%
set /a s1=(1%h1%*3600 + 1%time:~3,2%*60 + 1%time:~6,2%) - 111000

:: Cria a pasta de destino se não houver
if not exist dist (
    echo Criando pasta dist...
    mkdir dist
)
cd src

:: Compila para Windows
set h=%time:~0,2%
if "%h:~0,1%"==" " set h=0%h:~1,1%
echo %h%:%time:~3,5% - Compilando para Windows...
set GOOS=windows
set GOARCH=amd64
go build -o ..\dist\argo.exe
echo.

:: Compila para Linux
set h=%time:~0,2%
if "%h:~0,1%"==" " set h=0%h:~1,1%
echo %h%:%time:~3,5% - Compilando para Linux...
set GOOS=linux
set GOARCH=amd64
go build -o ..\dist\argo
echo.

:: Marca o fim do script
set h2=%time:~0,2%
if "%h2:~0,1%"==" " set h2=0%h2:~1,1%
set /a s2=(1%h2%*3600 + 1%time:~3,2%*60 + 1%time:~6,2%) - 111000

:: Calcula a duração
set /a duracao=%s2% - %s1%

:: Debug
set h=%time:~0,2%
if "%h:~0,1%"==" " set h=0%h:~1,1%
echo %h%:%time:~3,5% - Tempo total: %duracao% segs

echo.

cd..
