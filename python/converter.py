#This program converts a txt version on the bible into JSON so it is more readable
import json 

my_dict = {}

bible = {
    "Middleware": ("name", "release"),
    "System": ("name", "tag"),
    "Application": ("domain", "host", "user"),
    "Utility": ("domain", "health", "version"),
}

with open("test.txt") as file:
    # read through the input file, removing trailing newlines
    lines = [l.strip() for l in file.readlines()]

    # iterate through all the lines in the file
    for line in lines:
        # get the details from each line
        env, key, *field_values = line.split(",")
        # get the predefined keys from bible
        field_keys = bible[key]

        # create the 4 keys if the ENV does not exist
        # with defaults as empty list
        if env not in my_dict:
            my_dict[env] = {k: [] for k in bible}

        # append each new item to the existing list
        # zip is used to iterate through the keys and values to match them up
        my_dict[env][key].append({f_k: f_v for f_k, f_v in zip(field_keys, field_values)})


json_out = {"ENV": my_dict}
json.dump(json_out, "kjv.json")