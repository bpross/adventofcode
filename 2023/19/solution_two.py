import re


class PartsRatings:
    def __init__(self, parts_rating):
        ratings = re.search(
            r"x=(\d{1,4}),m=(\d{1,4}),a=(\d{1,4}),s=(\d{1,4})", parts_rating)
        self.x = int(ratings.group(1))
        self.m = int(ratings.group(2))
        self.a = int(ratings.group(3))
        self.s = int(ratings.group(4))

    def sum(self):
        return self.x + self.m + self.a + self.s

    def __repr__(self):
        return f"PartsRatings(x={self.x}, m={self.m}, a={self.a}, s={self.s})"


class Rule:
    def __init__(self, to_workflow, operator=None, value=None, field=None):
        self.to_workflow = to_workflow
        self.operator = operator
        self.value = value
        self.field = field

    def apply(self, parts_rating):
        if self.operator == '<':
            if getattr(parts_rating, self.field) < self.value:
                return self.to_workflow
            else:
                return None
        elif self.operator == '>':
            if getattr(parts_rating, self.field) > self.value:
                return self.to_workflow
            else:
                return None
        else:
            return self.to_workflow

    def __repr__(self):
        return f"Rule(to_workflow={self.to_workflow}, operator={self.operator}, value={self.value}, field={self.field})"


class Workflow:
    def __init__(self, label, rules):
        self.label = label
        self.rules = rules

    def __repr__(self):
        return f"Workflow(label={self.label}, rules={self.rules})"


def parse_workflow(workflow):
    label_and_rules = re.search(r"(\w{1,3}){(.+)}", workflow)
    label = label_and_rules.group(1)
    rules_group = label_and_rules.group(2).split(',')
    rules = []
    for rule in rules_group:
        rule_type = re.search(r"(\w{1})(<|>)(\d{1,4}):(\w{1,3})", rule)
        if rule_type:
            r = Rule(rule_type.group(4), rule_type.group(2), value=int(
                rule_type.group(3)), field=rule_type.group(1))
            rules.append(r)
        else:
            r = Rule(rule)
            rules.append(r)

    return label, rules

workflows = []
parts_ratings = []

workflows_parsed = False

with open('sample.txt') as f:
    for line in f:
        if not workflows_parsed:
            if line == '\n':
                workflows_parsed = True
                continue
            workflows.append(line.strip())
        else:
            parts_ratings.append(line.strip())

parts_ratings = [PartsRatings(rating) for rating in parts_ratings]
wflows = []
for workflow in workflows:
    label, rules = parse_workflow(workflow)
    for i in range(len(rules)):
        new_label = label 
        if i != 0:
            new_label = "{}_{}".format(label, i)
        new_rules = [rules[i]]
        if i + 1 < len(rules):
            next_label = "{}_{}".format(label, i + 1)
            new_rules.append(Rule(next_label))

        wflows.append(Workflow(new_label, new_rules))

w = {}
for workflow in wflows:
    w[workflow.label] = workflow

start = "in"
accepted = []
rejected = []

for rating in parts_ratings:
    current = start
    not_matched = True
    while not_matched:
        for rule in w[current].rules:
            next = rule.apply(rating)
            if next == "R":
                rejected.append(rating)
                not_matched = False
                break
            elif next == "A":
                not_matched = False
                accepted.append(rating)
                break
            elif next:
                current = next
                break

total = 0
for rating in accepted:
    total += rating.sum()

print(total)
