#!/usr/bin/env bash
# comment next line up if you are doing vagrant up separately (e.g. from a Windows machine)
vagrant up --provider=virtualbox
# add any playbook variables not in the roles as extra-vars below
ansible-playbook -i inventory -vvvv --extra-vars='{ "var_key":"var_value" }' site.yml
