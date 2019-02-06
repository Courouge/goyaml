The goal of this project is to transform Ansible projects with mixte YAML & Ansible syntax to pure YAML.

* Mixte Ansible/YAML syntax
```Ansible
  - name: insert trololo
    lineinfile: dest=/etc/trololo state=present regexp="{{ trololo }}" insertafter="trololo" line="trololo"
    notify: restart trololo
```
* Pure YAML syntax
```Ansible
  - name: insert trololo
    lineinfile:
      dest: /etc/trololo
      state: present
      regexp: "{{ trololo }}"
      insertafter: "trololo"
      line: "trololo"
    notify: restart trololo
```

* How it's works ?
```
go run *.go

```

* Bugs

```
- name: space in args module
  lineinfile: dest=/etc/trololo state=present regexp="{{ trololo_ trololo }}" insertafter="tro lolo " line="tro lo lo"
```

* Work with modules list in modules.txt
