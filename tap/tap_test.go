package tap

var buttonsNTabsOK = []byte(`buttons :
    - id : b1
      label : B 1
      panels :
        - id : not-ready
          note : some note
          markup : <p>b1 not-ready</p>
          myjs : my javascript not ready

        - id : ready
            tabs :
                - id : t1
                  label : T1
                  note: some note
                - id : t2
                  label : T2
                  note : some note
                - id : t3
                  label : T3
                  note : some note
    - id : b2
      label: B 2
      panels :
          - id : not-ready
            note : some note
          - id : ready
            buttons :
                  - id : b3
                    label : B3
                    note : some note
                  - id : b4
                    label : B4
                    note : some note
                  - id : b5
                    label : B5
                    note : some note
`)
