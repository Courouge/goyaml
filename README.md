The goal of this project is to transform your ansible project with mixte YAML & Ansible syntax to pure YAML.

* Mixte Ansible/YAML syntax
```Ansible
  - name: insert iptables rule
    lineinfile: dest=/etc/sysconfig/iptables state=present regexp="{{ mysql_port }}" insertafter="trololo " line="trololo"
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
- name: space in args
  lineinfile: dest=/etc/sysconfig/iptables state=present regexp="{{ mysql_port }}" insertafter="trololo " line="tro lo lo"
```
