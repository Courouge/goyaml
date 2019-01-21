The goal of this project is to transform your ansible project with mixte YAML & Ansible syntax to pure YAML.

Example:
(Mixte Ansible/YAML syntax)
  - name: insert iptables rule
    lineinfile: dest=/etc/sysconfig/iptables state=present regexp="{{ mysql_port }}" insertafter="^:OUTPUT " line="-A INPUT -p tcp --dport {{ mysql_port }} -j ACCEPT"
    notify: restart iptables
    
(pure YAML syntax)
  - name: insert iptables rule
    lineinfile: 
      dest: /etc/sysconfig/iptables
      state: present 
      regexp: "{{ mysql_port }}" 
      insertafter: "^:OUTPUT " line="-A INPUT -p tcp --dport {{ mysql_port }} -j ACCEPT"
    notify: restart iptables
