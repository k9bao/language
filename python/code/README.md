# 说明

- package:sensitive
  - module:sensitive_codec
    - python -m sensitive.sensitive_codec --help
    - python -m sensitive.sensitive_codec decode-file
    - python -m sensitive.sensitive_codec encode-json
    - python -m sensitive.sensitive_codec check-file
  - module:sensitive_replace
    - python -m sensitive.sensitive_replace --help
    - python -m sensitive.sensitive_replace replace-root --dir=testdata
    - python -m sensitive.sensitive_replace replace-root --dir=testdata --enc=0
    - python -m sensitive.sensitive_replace replace-root --dir=testdata --enc=1

- package:fs
  - module:replace
    - python -m fs.dirUtil
    - python -m fs.file_replace
