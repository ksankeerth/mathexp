# matheXP

matheXP is a small library that helps to create shareable complex expressions using JSON.

## Features

Read complex expressions in JSON format and evaluate against the provided data.

## Usecases

In most cases, We usually have static expressions and equations in code. In some cases, the end-user needs to create expressions with conditions and evaluate them with data while the program/server is running. The matheXp is developed for such use cases.

The matheXp uses a JSON to represent expressions and conditions. So the expressions can be created using a User Interface simply.

For example, if we want to calculate the electricity bill with the below conditions, We can have if and else conditions in the code and evaluate the cost simply. But if an end-user wants to create such an equation dynamically when a program is running, the matheXp will be a good choice.

Note: For simplicity, the electricity bill calculation is shown as an example. Actually, this library was created for supporting calculations of Green House Gas (GHG) emissions from many varying sources


| kWh | kWh   | Cost |
| :---- | ------- | :----- |
| 0   | 100   | 5.0  |
| 100 | 200   | 7.0  |
| 200 | above | 10.0 |

In matheXp, the equations for above can be represented as below.

```
{
    "vars": [
        {
            "type": "const",
            "sym": "elec_prc_under_100kwh",
            "name": "Electricy Price per kWH if the usage under 100KWH",
            "value": 5.0
        },
        {
            "type": "const",
            "sym": "elec_prc_bw_100kwh_and_200kwh",
            "name": "Electricy Price per kWH if the usage between 100kWH and 200kwh",
            "value": 7.0
        },
        {
            "type": "const",
            "sym": "elec_prc_above_200kwh",
            "name": "Electricy Price per kWH if the usage is over 200kwh",
            "value": 10.0
        },
        {
            "type": "in",
            "sym": "elec_consumption",
            "name": "Electricity consumption"
        },
        {
            "type": "out",
            "sym": "elec_cost",
            "name": "Cost of Electricity"
        }
    ],
    "expressions": [],
    "subconditiongroups": [
        {
            "cond": {
                "op": "lt",
                "v1": "elec_consumption",
                "v2": 100
            },
            "subconditiongroups": [],
            "expressions": [
                {
                    "op": "mul",
                    "v1": "elec_consumption",
                    "v2": "elec_prc_under_100kwh",
                    "out": "elec_cost"
                }
            ],
            "vars": []
        },
        {
            "cond": {
                "op": "and",
                "v1": {
                    "op": "gte",
                    "v1": "elec_consumption",
                    "v2": 100
                },
                "v2": {
                    "op": "lt",
                    "v1": "elec_consumption",
                    "v2": 200
                }
            },
            "subconditiongroups": [],
            "expressions": [
                {
                    "op": "mul",
                    "v1": "elec_consumption",
                    "v2": "elec_prc_bw_100kwh_and_200kwh",
                    "out": "elec_cost"
                }
            ],
            "vars": []
        },
        {
            "cond": {
                "op": "gt",
                "v1": "elec_consumption",
                "v2": 200
            },
            "subconditiongroups": [],
            "expressions": [
                {
                    "op": "mul",
                    "v1": "elec_consumption",
                    "v2": "elec_prc_above_200kwh",
                    "out": "elec_cost"
                }
            ],
            "vars": []
        }
    ],
    "cond": null
}
```

### Supported Conditional Operators


| Symbol | Description            |
| :------- | :----------------------- |
| lt     | Lesser than            |
| lte    | Lesser than and equal  |
| gt     | Greater than           |
| gte    | Greater than and equal |
| eq     | Equals to              |
| neq    | Not equals to          |
| and    | Logical and            |
| or     | Logical or             |

### Supported Mathematical Operators


| Symbol | Description     |
| :------- | ----------------- |
| add    | Addition        |
| sub    | Subtraction     |
| mul    | Multifilication |
| div    | Division        |
