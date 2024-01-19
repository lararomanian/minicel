# MiniCel - Basic String Parser / Excel Implementation in Go

MiniCel is a simple string parser and Excel-like implementation written in Go, inspired by Tsoding's project "minicel."

## Overview

MiniCel allows you to create and manipulate data in a format similar to Excel. It supports basic operations such as cell referencing and simple formulas.

## How it Works

The structure for the CSV files used by MiniCel follows a basic Excel-like format:

# Features

 **Cell Referencing**: Refer to cells in your formulas using standard Excel-style references (e.g., A1, B2).

**Basic Formulas**: Perform basic calculations using built-in functions such as SUM, AVERAGE, etc.

**CSV File Support**: Save and load your MiniCel sheets in CSV format.

# Examples

## Simple Calculation

| A | B | C       |
|---|---|---------|
| 1 | 2 | =SUM(A1:B1) |

### Result

| A | B | C |
|---|---|---|
| 1 | 2 | 3 |

## Referencing Cells

| A | B | C       |
|---|---|---------|
| 1 | 2 | =A1 + B1 |

### Result

| A | B | C |
|---|---|---|
| 1 | 2 | 3 |




