The goal of this project is to transform Ansible projects with mixte YAML & Ansible syntax to pure YAML.

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
- name: space in args
  lineinfile: dest=/etc/sysconfig/iptables state=present regexp="{{ mysql _port }}" insertafter="tro lolo " line="tro lo lo"
```

* Work only with this modules list
  apt
  service
  wait_for
  file
  line
  lineinfile
  package
  stat
  sysctl
  uri
  user
  systemd
  copy
  replace
  alternatives
  stat
  file
  template
