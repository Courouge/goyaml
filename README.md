The goal of this project is to transform Ansible projects with mixte YAML & Ansible syntax to pure YAML.

* How it's works ?
```
go run *.go

```
* Mixte Ansible/YAML syntax
```Ansible
  - name: insert iptables rule
    lineinfile: dest=/etc/sysconfig/iptables state=present regexp="{{ mysql_port }}" insertafter="trololo" line="trololo"
    notify: restart iptables
```
* Pure YAML syntax
```Ansible
  - name: insert iptables rule
    lineinfile: 
      dest: /etc/sysconfig/iptables
      state: present 
      regexp: "{{ mysql_port }}" 
      insertafter: "trololo"
      line: "trololo"
    notify: restart iptables
```

* Bugs

```
- name: space in args module
  lineinfile: dest=/etc/sysconfig/iptables state=present regexp="{{ mysql _port }}" insertafter="tro lolo " line="tro lo lo"
```

* Work with modules list in modules.txt
