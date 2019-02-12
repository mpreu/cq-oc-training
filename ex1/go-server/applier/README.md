ansible-galaxy install -r requirements.yml --roles-path=roles --f
ansible-playbook apply.yml -i inventory/ -e target=seed-hosts -e provision=true

