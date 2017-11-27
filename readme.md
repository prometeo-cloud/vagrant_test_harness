## vagrant_test_narness

A simple test harness project to facilitate testing Ansible roles on Vagrant boxes.

### Using the project

1. Update the requirements.yml to include the role you want to test and its dependent roles if any.
2. Run **sh download** to recreate the roles folder locally with the roles in the requirements.yml file.
3. Ensure you have a **rhel74_sub** local vagrant box, if not see [here](https://github.com/prometeo-cloud/vagrant_create_subscribed_box) for information on how to create one.
4. Only if you need it, update the **site.yml** file. The test_role in the file should not need changing as it is deployed by ansible-galaxy from the requirements.yml file.
5. Only if you are spinning up a host separately (update the **go** file)
6. Run **sh go**