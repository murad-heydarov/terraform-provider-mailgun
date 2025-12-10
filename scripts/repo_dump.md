

===== FILE: ./GNUmakefile =====

TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=mailgun

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m -count=1

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

vendor-status:
	@govendor status

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc vet fmt fmtcheck errcheck vendor-status test-compile website website-test



===== FILE: ./go.mod =====

module github.com/terraform-providers/terraform-provider-mailgun

go 1.24.1

require (
	github.com/hashicorp/go-uuid v1.0.3
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.37.0
	github.com/mailgun/mailgun-go/v5 v5.6.2
)

require (
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/cloudflare/circl v1.6.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.5.0 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.6.3 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/hc-install v0.9.2 // indirect
	github.com/hashicorp/hcl/v2 v2.23.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.23.0 // indirect
	github.com/hashicorp/terraform-json v0.25.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.27.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.5 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kisielk/errcheck v1.9.0 // indirect
	github.com/mailgun/errors v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oapi-codegen/runtime v1.1.2 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/zclconf/go-cty v1.16.2 // indirect
	golang.org/x/crypto v0.42.0 // indirect
	golang.org/x/mod v0.28.0 // indirect
	golang.org/x/net v0.44.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/text v0.29.0 // indirect
	golang.org/x/tools v0.37.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.72.1 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)


===== FILE: ./LICENSE =====

Mozilla Public License Version 2.0
==================================

1. Definitions
--------------

1.1. "Contributor"
    means each individual or legal entity that creates, contributes to
    the creation of, or owns Covered Software.

1.2. "Contributor Version"
    means the combination of the Contributions of others (if any) used
    by a Contributor and that particular Contributor's Contribution.

1.3. "Contribution"
    means Covered Software of a particular Contributor.

1.4. "Covered Software"
    means Source Code Form to which the initial Contributor has attached
    the notice in Exhibit A, the Executable Form of such Source Code
    Form, and Modifications of such Source Code Form, in each case
    including portions thereof.

1.5. "Incompatible With Secondary Licenses"
    means

    (a) that the initial Contributor has attached the notice described
        in Exhibit B to the Covered Software; or

    (b) that the Covered Software was made available under the terms of
        version 1.1 or earlier of the License, but not also under the
        terms of a Secondary License.

1.6. "Executable Form"
    means any form of the work other than Source Code Form.

1.7. "Larger Work"
    means a work that combines Covered Software with other material, in
    a separate file or files, that is not Covered Software.

1.8. "License"
    means this document.

1.9. "Licensable"
    means having the right to grant, to the maximum extent possible,
    whether at the time of the initial grant or subsequently, any and
    all of the rights conveyed by this License.

1.10. "Modifications"
    means any of the following:

    (a) any file in Source Code Form that results from an addition to,
        deletion from, or modification of the contents of Covered
        Software; or

    (b) any new file in Source Code Form that contains any Covered
        Software.

1.11. "Patent Claims" of a Contributor
    means any patent claim(s), including without limitation, method,
    process, and apparatus claims, in any patent Licensable by such
    Contributor that would be infringed, but for the grant of the
    License, by the making, using, selling, offering for sale, having
    made, import, or transfer of either its Contributions or its
    Contributor Version.

1.12. "Secondary License"
    means either the GNU General Public License, Version 2.0, the GNU
    Lesser General Public License, Version 2.1, the GNU Affero General
    Public License, Version 3.0, or any later versions of those
    licenses.

1.13. "Source Code Form"
    means the form of the work preferred for making modifications.

1.14. "You" (or "Your")
    means an individual or a legal entity exercising rights under this
    License. For legal entities, "You" includes any entity that
    controls, is controlled by, or is under common control with You. For
    purposes of this definition, "control" means (a) the power, direct
    or indirect, to cause the direction or management of such entity,
    whether by contract or otherwise, or (b) ownership of more than
    fifty percent (50%) of the outstanding shares or beneficial
    ownership of such entity.

2. License Grants and Conditions
--------------------------------

2.1. Grants

Each Contributor hereby grants You a world-wide, royalty-free,
non-exclusive license:

(a) under intellectual property rights (other than patent or trademark)
    Licensable by such Contributor to use, reproduce, make available,
    modify, display, perform, distribute, and otherwise exploit its
    Contributions, either on an unmodified basis, with Modifications, or
    as part of a Larger Work; and

(b) under Patent Claims of such Contributor to make, use, sell, offer
    for sale, have made, import, and otherwise transfer either its
    Contributions or its Contributor Version.

2.2. Effective Date

The licenses granted in Section 2.1 with respect to any Contribution
become effective for each Contribution on the date the Contributor first
distributes such Contribution.

2.3. Limitations on Grant Scope

The licenses granted in this Section 2 are the only rights granted under
this License. No additional rights or licenses will be implied from the
distribution or licensing of Covered Software under this License.
Notwithstanding Section 2.1(b) above, no patent license is granted by a
Contributor:

(a) for any code that a Contributor has removed from Covered Software;
    or

(b) for infringements caused by: (i) Your and any other third party's
    modifications of Covered Software, or (ii) the combination of its
    Contributions with other software (except as part of its Contributor
    Version); or

(c) under Patent Claims infringed by Covered Software in the absence of
    its Contributions.

This License does not grant any rights in the trademarks, service marks,
or logos of any Contributor (except as may be necessary to comply with
the notice requirements in Section 3.4).

2.4. Subsequent Licenses

No Contributor makes additional grants as a result of Your choice to
distribute the Covered Software under a subsequent version of this
License (see Section 10.2) or under the terms of a Secondary License (if
permitted under the terms of Section 3.3).

2.5. Representation

Each Contributor represents that the Contributor believes its
Contributions are its original creation(s) or it has sufficient rights
to grant the rights to its Contributions conveyed by this License.

2.6. Fair Use

This License is not intended to limit any rights You have under
applicable copyright doctrines of fair use, fair dealing, or other
equivalents.

2.7. Conditions

Sections 3.1, 3.2, 3.3, and 3.4 are conditions of the licenses granted
in Section 2.1.

3. Responsibilities
-------------------

3.1. Distribution of Source Form

All distribution of Covered Software in Source Code Form, including any
Modifications that You create or to which You contribute, must be under
the terms of this License. You must inform recipients that the Source
Code Form of the Covered Software is governed by the terms of this
License, and how they can obtain a copy of this License. You may not
attempt to alter or restrict the recipients' rights in the Source Code
Form.

3.2. Distribution of Executable Form

If You distribute Covered Software in Executable Form then:

(a) such Covered Software must also be made available in Source Code
    Form, as described in Section 3.1, and You must inform recipients of
    the Executable Form how they can obtain a copy of such Source Code
    Form by reasonable means in a timely manner, at a charge no more
    than the cost of distribution to the recipient; and

(b) You may distribute such Executable Form under the terms of this
    License, or sublicense it under different terms, provided that the
    license for the Executable Form does not attempt to limit or alter
    the recipients' rights in the Source Code Form under this License.

3.3. Distribution of a Larger Work

You may create and distribute a Larger Work under terms of Your choice,
provided that You also comply with the requirements of this License for
the Covered Software. If the Larger Work is a combination of Covered
Software with a work governed by one or more Secondary Licenses, and the
Covered Software is not Incompatible With Secondary Licenses, this
License permits You to additionally distribute such Covered Software
under the terms of such Secondary License(s), so that the recipient of
the Larger Work may, at their option, further distribute the Covered
Software under the terms of either this License or such Secondary
License(s).

3.4. Notices

You may not remove or alter the substance of any license notices
(including copyright notices, patent notices, disclaimers of warranty,
or limitations of liability) contained within the Source Code Form of
the Covered Software, except that You may alter any license notices to
the extent required to remedy known factual inaccuracies.

3.5. Application of Additional Terms

You may choose to offer, and to charge a fee for, warranty, support,
indemnity or liability obligations to one or more recipients of Covered
Software. However, You may do so only on Your own behalf, and not on
behalf of any Contributor. You must make it absolutely clear that any
such warranty, support, indemnity, or liability obligation is offered by
You alone, and You hereby agree to indemnify every Contributor for any
liability incurred by such Contributor as a result of warranty, support,
indemnity or liability terms You offer. You may include additional
disclaimers of warranty and limitations of liability specific to any
jurisdiction.

4. Inability to Comply Due to Statute or Regulation
---------------------------------------------------

If it is impossible for You to comply with any of the terms of this
License with respect to some or all of the Covered Software due to
statute, judicial order, or regulation then You must: (a) comply with
the terms of this License to the maximum extent possible; and (b)
describe the limitations and the code they affect. Such description must
be placed in a text file included with all distributions of the Covered
Software under this License. Except to the extent prohibited by statute
or regulation, such description must be sufficiently detailed for a
recipient of ordinary skill to be able to understand it.

5. Termination
--------------

5.1. The rights granted under this License will terminate automatically
if You fail to comply with any of its terms. However, if You become
compliant, then the rights granted under this License from a particular
Contributor are reinstated (a) provisionally, unless and until such
Contributor explicitly and finally terminates Your grants, and (b) on an
ongoing basis, if such Contributor fails to notify You of the
non-compliance by some reasonable means prior to 60 days after You have
come back into compliance. Moreover, Your grants from a particular
Contributor are reinstated on an ongoing basis if such Contributor
notifies You of the non-compliance by some reasonable means, this is the
first time You have received notice of non-compliance with this License
from such Contributor, and You become compliant prior to 30 days after
Your receipt of the notice.

5.2. If You initiate litigation against any entity by asserting a patent
infringement claim (excluding declaratory judgment actions,
counter-claims, and cross-claims) alleging that a Contributor Version
directly or indirectly infringes any patent, then the rights granted to
You by any and all Contributors for the Covered Software under Section
2.1 of this License shall terminate.

5.3. In the event of termination under Sections 5.1 or 5.2 above, all
end user license agreements (excluding distributors and resellers) which
have been validly granted by You or Your distributors under this License
prior to termination shall survive termination.

************************************************************************
*                                                                      *
*  6. Disclaimer of Warranty                                           *
*  -------------------------                                           *
*                                                                      *
*  Covered Software is provided under this License on an "as is"       *
*  basis, without warranty of any kind, either expressed, implied, or  *
*  statutory, including, without limitation, warranties that the       *
*  Covered Software is free of defects, merchantable, fit for a        *
*  particular purpose or non-infringing. The entire risk as to the     *
*  quality and performance of the Covered Software is with You.        *
*  Should any Covered Software prove defective in any respect, You     *
*  (not any Contributor) assume the cost of any necessary servicing,   *
*  repair, or correction. This disclaimer of warranty constitutes an   *
*  essential part of this License. No use of any Covered Software is   *
*  authorized under this License except under this disclaimer.         *
*                                                                      *
************************************************************************

************************************************************************
*                                                                      *
*  7. Limitation of Liability                                          *
*  --------------------------                                          *
*                                                                      *
*  Under no circumstances and under no legal theory, whether tort      *
*  (including negligence), contract, or otherwise, shall any           *
*  Contributor, or anyone who distributes Covered Software as          *
*  permitted above, be liable to You for any direct, indirect,         *
*  special, incidental, or consequential damages of any character      *
*  including, without limitation, damages for lost profits, loss of    *
*  goodwill, work stoppage, computer failure or malfunction, or any    *
*  and all other commercial damages or losses, even if such party      *
*  shall have been informed of the possibility of such damages. This   *
*  limitation of liability shall not apply to liability for death or   *
*  personal injury resulting from such party's negligence to the       *
*  extent applicable law prohibits such limitation. Some               *
*  jurisdictions do not allow the exclusion or limitation of           *
*  incidental or consequential damages, so this exclusion and          *
*  limitation may not apply to You.                                    *
*                                                                      *
************************************************************************

8. Litigation
-------------

Any litigation relating to this License may be brought only in the
courts of a jurisdiction where the defendant maintains its principal
place of business and such litigation shall be governed by laws of that
jurisdiction, without reference to its conflict-of-law provisions.
Nothing in this Section shall prevent a party's ability to bring
cross-claims or counter-claims.

9. Miscellaneous
----------------

This License represents the complete agreement concerning the subject
matter hereof. If any provision of this License is held to be
unenforceable, such provision shall be reformed only to the extent
necessary to make it enforceable. Any law or regulation which provides
that the language of a contract shall be construed against the drafter
shall not be used to construe this License against a Contributor.

10. Versions of the License
---------------------------

10.1. New Versions

Mozilla Foundation is the license steward. Except as provided in Section
10.3, no one other than the license steward has the right to modify or
publish new versions of this License. Each version will be given a
distinguishing version number.

10.2. Effect of New Versions

You may distribute the Covered Software under the terms of the version
of the License under which You originally received the Covered Software,
or under the terms of any subsequent version published by the license
steward.

10.3. Modified Versions

If you create software not governed by this License, and you want to
create a new license for such software, you may create and use a
modified version of this License if you rename the license and remove
any references to the name of the license steward (except to note that
such modified license differs from this License).

10.4. Distributing Source Code Form that is Incompatible With Secondary
Licenses

If You choose to distribute Source Code Form that is Incompatible With
Secondary Licenses under the terms of this version of the License, the
notice described in Exhibit B of this License must be attached.

Exhibit A - Source Code Form License Notice
-------------------------------------------

  This Source Code Form is subject to the terms of the Mozilla Public
  License, v. 2.0. If a copy of the MPL was not distributed with this
  file, You can obtain one at http://mozilla.org/MPL/2.0/.

If it is not possible or desirable to put the notice in a particular
file, then You may include the notice in a location (such as a LICENSE
file in a relevant directory) where a recipient would be likely to look
for such a notice.

You may add additional accurate notices of copyright ownership.

Exhibit B - "Incompatible With Secondary Licenses" Notice
---------------------------------------------------------

  This Source Code Form is "Incompatible With Secondary Licenses", as
  defined by the Mozilla Public License, v. 2.0.


===== FILE: ./go.sum =====

dario.cat/mergo v1.0.0 h1:AGCNq9Evsj31mOgNPcLyXc+4PNABt905YmuqPYYpBWk=
dario.cat/mergo v1.0.0/go.mod h1:uNxQE+84aUszobStD9th8a29P2fMDhsBdgRYvZOxGmk=
github.com/Microsoft/go-winio v0.6.2 h1:F2VQgta7ecxGYO8k3ZZz3RS8fVIXVxONVUPlNERoyfY=
github.com/Microsoft/go-winio v0.6.2/go.mod h1:yd8OoFMLzJbo9gZq8j5qaps8bJ9aShtEA8Ipt1oGCvU=
github.com/ProtonMail/go-crypto v1.1.6 h1:ZcV+Ropw6Qn0AX9brlQLAUXfqLBc7Bl+f/DmNxpLfdw=
github.com/ProtonMail/go-crypto v1.1.6/go.mod h1:rA3QumHc/FZ8pAHreoekgiAbzpNsfQAosU5td4SnOrE=
github.com/agext/levenshtein v1.2.2 h1:0S/Yg6LYmFJ5stwQeRp6EeOcCbj7xiqQSdNelsXvaqE=
github.com/agext/levenshtein v1.2.2/go.mod h1:JEDfjyjHDjOF/1e4FlBE/PkbqA9OfWu2ki2W0IB5558=
github.com/apparentlymart/go-textseg/v12 v12.0.0/go.mod h1:S/4uRK2UtaQttw1GenVJEynmyUenKwP++x/+DdGV/Ec=
github.com/apparentlymart/go-textseg/v15 v15.0.0 h1:uYvfpb3DyLSCGWnctWKGj857c6ew1u1fNQOlOtuGxQY=
github.com/apparentlymart/go-textseg/v15 v15.0.0/go.mod h1:K8XmNZdhEBkdlyDdvbmmsvpAG721bKi0joRfFdHIWJ4=
github.com/bufbuild/protocompile v0.4.0 h1:LbFKd2XowZvQ/kajzguUp2DC9UEIQhIq77fZZlaQsNA=
github.com/bufbuild/protocompile v0.4.0/go.mod h1:3v93+mbWn/v3xzN+31nwkJfrEpAUwp+BagBSZWx+TP8=
github.com/cloudflare/circl v1.6.0 h1:cr5JKic4HI+LkINy2lg3W2jF8sHCVTBncJr5gIIq7qk=
github.com/cloudflare/circl v1.6.0/go.mod h1:uddAzsPgqdMAYatqJ0lsjX1oECcQLIlRpzZh3pJrofs=
github.com/cyphar/filepath-securejoin v0.4.1 h1:JyxxyPEaktOD+GAnqIqTf9A8tHyAG22rowi7HkoSU1s=
github.com/cyphar/filepath-securejoin v0.4.1/go.mod h1:Sdj7gXlvMcPZsbhwhQ33GguGLDGQL7h7bg04C/+u9jI=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/emirpasic/gods v1.18.1 h1:FXtiHYKDGKCW2KzwZKx0iC0PQmdlorYgdFG9jPXJ1Bc=
github.com/emirpasic/gods v1.18.1/go.mod h1:8tpGGwCnJ5H4r6BWwaV6OrWmMoPhUl5jm/FMNAnJvWQ=
github.com/fatih/color v1.13.0/go.mod h1:kLAiJbzzSOZDVNGyDpeOxJ47H46qBXwg5ILebYFFOfk=
github.com/fatih/color v1.16.0 h1:zmkK9Ngbjj+K0yRhTVONQh1p/HknKYSlNT+vZCzyokM=
github.com/fatih/color v1.16.0/go.mod h1:fL2Sau1YI5c0pdGEVCbKQbLXB6edEj1ZgiY4NijnWvE=
github.com/go-chi/chi/v5 v5.2.2 h1:CMwsvRVTbXVytCk1Wd72Zy1LAsAh9GxMmSNWLHCG618=
github.com/go-chi/chi/v5 v5.2.2/go.mod h1:L2yAIGWB3H+phAw1NxKwWM+7eUH/lU8pOMm5hHcoops=
github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 h1:+zs/tPmkDkHx3U66DAb0lQFJrpS6731Oaa12ikc+DiI=
github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376/go.mod h1:an3vInlBmSxCcxctByoQdvwPiA7DTK7jaaFDBTtu0ic=
github.com/go-git/go-billy/v5 v5.6.2 h1:6Q86EsPXMa7c3YZ3aLAQsMA0VlWmy43r6FHqa/UNbRM=
github.com/go-git/go-billy/v5 v5.6.2/go.mod h1:rcFC2rAsp/erv7CMz9GczHcuD0D32fWzH+MJAU+jaUU=
github.com/go-git/go-git/v5 v5.14.0 h1:/MD3lCrGjCen5WfEAzKg00MJJffKhC8gzS80ycmCi60=
github.com/go-git/go-git/v5 v5.14.0/go.mod h1:Z5Xhoia5PcWA3NF8vRLURn9E5FRhSl7dGj9ItW3Wk5k=
github.com/go-logr/logr v1.4.2 h1:6pFjapn8bFcIbiKo3XT4j/BhANplGihG6tvd+8rYgrY=
github.com/go-logr/logr v1.4.2/go.mod h1:9T104GzyrTigFIr8wt5mBrctHMim0Nb2HLGrmQ40KvY=
github.com/go-logr/stdr v1.2.2 h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=
github.com/go-logr/stdr v1.2.2/go.mod h1:mMo/vtBO5dYbehREoey6XUKy/eSumjCCveDpRre4VKE=
github.com/go-test/deep v1.0.3 h1:ZrJSEWsXzPOxaZnFteGEfooLba+ju3FYIbOrS+rQd68=
github.com/go-test/deep v1.0.3/go.mod h1:wGDj63lr65AM2AQyKZd/NYHGb0R+1RLqB8NKt3aSFNA=
github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 h1:f+oWsMOmNPc8JmEHVZIycC7hBoQxHH9pNKQORJNozsQ=
github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8/go.mod h1:wcDNUvekVysuuOpQKo3191zZyTpiI6se1N1ULghS0sw=
github.com/golang/protobuf v1.1.0/go.mod h1:6lQm79b+lXiMfvg/cZm0SGofjICqVBUtrP5yJMmIC1U=
github.com/golang/protobuf v1.5.0/go.mod h1:FsONVRAS9T7sI+LIUmWTfcYkHO4aIWwzhcaSAoJOfIk=
github.com/golang/protobuf v1.5.2/go.mod h1:XVQd3VNwM+JqD3oG2Ue2ip4fOMUkwXdXDdiuN0vRsmY=
github.com/golang/protobuf v1.5.4 h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=
github.com/golang/protobuf v1.5.4/go.mod h1:lnTiLA8Wa4RWRcIUkrtSVa5nRhsEGBg48fD6rSs7xps=
github.com/google/go-cmp v0.3.1/go.mod h1:8QqcDgzrUqlUb/G2PQTWiueGozuR1884gddMywk6iLU=
github.com/google/go-cmp v0.5.5/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
github.com/google/go-cmp v0.7.0 h1:wk8382ETsv4JYUZwIsn6YpYiWiBsYLSJiTsyBybVuN8=
github.com/google/go-cmp v0.7.0/go.mod h1:pXiqmnSA92OHEEa9HXL2W4E7lf9JzCmGVUdgjX3N/iU=
github.com/google/gofuzz v1.0.0/go.mod h1:dBl0BpW6vV/+mYPU4Po3pmUjxk6FQPldtuIdl/M65Eg=
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/hashicorp/errwrap v1.0.0 h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=
github.com/hashicorp/errwrap v1.0.0/go.mod h1:YH+1FKiLXxHSkmPseP+kNlulaMuP3n2brvKWEqk/Jc4=
github.com/hashicorp/go-checkpoint v0.5.0 h1:MFYpPZCnQqQTE18jFwSII6eUQrD/oxMFp3mlgcqk5mU=
github.com/hashicorp/go-checkpoint v0.5.0/go.mod h1:7nfLNL10NsxqO4iWuW6tWW0HjZuDrwkBuEQsVcpCOgg=
github.com/hashicorp/go-cleanhttp v0.5.0/go.mod h1:JpRdi6/HCYpAwUzNwuwqhbovhLtngrth3wmdIIUrZ80=
github.com/hashicorp/go-cleanhttp v0.5.2 h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=
github.com/hashicorp/go-cleanhttp v0.5.2/go.mod h1:kO/YDlP8L1346E6Sodw+PrpBSV4/SoxCXGY6BqNFT48=
github.com/hashicorp/go-cty v1.5.0 h1:EkQ/v+dDNUqnuVpmS5fPqyY71NXVgT5gf32+57xY8g0=
github.com/hashicorp/go-cty v1.5.0/go.mod h1:lFUCG5kd8exDobgSfyj4ONE/dc822kiYMguVKdHGMLM=
github.com/hashicorp/go-hclog v1.6.3 h1:Qr2kF+eVWjTiYmU7Y31tYlP1h0q/X3Nl3tPGdaB11/k=
github.com/hashicorp/go-hclog v1.6.3/go.mod h1:W4Qnvbt70Wk/zYJryRzDRU/4r0kIg0PVHBcfoyhpF5M=
github.com/hashicorp/go-multierror v1.1.1 h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=
github.com/hashicorp/go-multierror v1.1.1/go.mod h1:iw975J/qwKPdAO1clOe2L8331t/9/fmwbPZ6JB6eMoM=
github.com/hashicorp/go-plugin v1.6.3 h1:xgHB+ZUSYeuJi96WtxEjzi23uh7YQpznjGh0U0UUrwg=
github.com/hashicorp/go-plugin v1.6.3/go.mod h1:MRobyh+Wc/nYy1V4KAXUiYfzxoYhs7V1mlH1Z7iY2h0=
github.com/hashicorp/go-retryablehttp v0.7.7 h1:C8hUCYzor8PIfXHa4UrZkU4VvK8o9ISHxT2Q8+VepXU=
github.com/hashicorp/go-retryablehttp v0.7.7/go.mod h1:pkQpWZeYWskR+D1tR2O5OcBFOxfA7DoAO6xtkuQnHTk=
github.com/hashicorp/go-uuid v1.0.0/go.mod h1:6SBZvOh/SIDV7/2o3Jml5SYk/TvGqwFJ/bN7x4byOro=
github.com/hashicorp/go-uuid v1.0.3 h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=
github.com/hashicorp/go-uuid v1.0.3/go.mod h1:6SBZvOh/SIDV7/2o3Jml5SYk/TvGqwFJ/bN7x4byOro=
github.com/hashicorp/go-version v1.7.0 h1:5tqGy27NaOTB8yJKUZELlFAS/LTKJkrmONwQKeRZfjY=
github.com/hashicorp/go-version v1.7.0/go.mod h1:fltr4n8CU8Ke44wwGCBoEymUuxUHl09ZGVZPK5anwXA=
github.com/hashicorp/hc-install v0.9.2 h1:v80EtNX4fCVHqzL9Lg/2xkp62bbvQMnvPQ0G+OmtO24=
github.com/hashicorp/hc-install v0.9.2/go.mod h1:XUqBQNnuT4RsxoxiM9ZaUk0NX8hi2h+Lb6/c0OZnC/I=
github.com/hashicorp/hcl/v2 v2.23.0 h1:Fphj1/gCylPxHutVSEOf2fBOh1VE4AuLV7+kbJf3qos=
github.com/hashicorp/hcl/v2 v2.23.0/go.mod h1:62ZYHrXgPoX8xBnzl8QzbWq4dyDsDtfCRgIq1rbJEvA=
github.com/hashicorp/logutils v1.0.0 h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=
github.com/hashicorp/logutils v1.0.0/go.mod h1:QIAnNjmIWmVIIkWDTG1z5v++HQmx9WQRO+LraFDTW64=
github.com/hashicorp/terraform-exec v0.23.0 h1:MUiBM1s0CNlRFsCLJuM5wXZrzA3MnPYEsiXmzATMW/I=
github.com/hashicorp/terraform-exec v0.23.0/go.mod h1:mA+qnx1R8eePycfwKkCRk3Wy65mwInvlpAeOwmA7vlY=
github.com/hashicorp/terraform-json v0.25.0 h1:rmNqc/CIfcWawGiwXmRuiXJKEiJu1ntGoxseG1hLhoQ=
github.com/hashicorp/terraform-json v0.25.0/go.mod h1:sMKS8fiRDX4rVlR6EJUMudg1WcanxCMoWwTLkgZP/vc=
github.com/hashicorp/terraform-plugin-go v0.27.0 h1:ujykws/fWIdsi6oTUT5Or4ukvEan4aN9lY+LOxVP8EE=
github.com/hashicorp/terraform-plugin-go v0.27.0/go.mod h1:FDa2Bb3uumkTGSkTFpWSOwWJDwA7bf3vdP3ltLDTH6o=
github.com/hashicorp/terraform-plugin-log v0.9.0 h1:i7hOA+vdAItN1/7UrfBqBwvYPQ9TFvymaRGZED3FCV0=
github.com/hashicorp/terraform-plugin-log v0.9.0/go.mod h1:rKL8egZQ/eXSyDqzLUuwUYLVdlYeamldAHSxjUFADow=
github.com/hashicorp/terraform-plugin-sdk/v2 v2.37.0 h1:NFPMacTrY/IdcIcnUB+7hsore1ZaRWU9cnB6jFoBnIM=
github.com/hashicorp/terraform-plugin-sdk/v2 v2.37.0/go.mod h1:QYmYnLfsosrxjCnGY1p9c7Zj6n9thnEE+7RObeYs3fA=
github.com/hashicorp/terraform-registry-address v0.2.5 h1:2GTftHqmUhVOeuu9CW3kwDkRe4pcBDq0uuK5VJngU1M=
github.com/hashicorp/terraform-registry-address v0.2.5/go.mod h1:PpzXWINwB5kuVS5CA7m1+eO2f1jKb5ZDIxrOPfpnGkg=
github.com/hashicorp/terraform-svchost v0.1.1 h1:EZZimZ1GxdqFRinZ1tpJwVxxt49xc/S52uzrw4x0jKQ=
github.com/hashicorp/terraform-svchost v0.1.1/go.mod h1:mNsjQfZyf/Jhz35v6/0LWcv26+X7JPS+buii2c9/ctc=
github.com/hashicorp/yamux v0.1.1 h1:yrQxtgseBDrq9Y652vSRDvsKCJKOUD+GzTS4Y0Y8pvE=
github.com/hashicorp/yamux v0.1.1/go.mod h1:CtWFDAQgb7dxtzFs4tWbplKIe2jSi3+5vKbgIO0SLnQ=
github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 h1:BQSFePA1RWJOlocH6Fxy8MmwDt+yVQYULKfN0RoTN8A=
github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99/go.mod h1:1lJo3i6rXxKeerYnT8Nvf0QmHCRC1n8sfWVwXF2Frvo=
github.com/jhump/protoreflect v1.15.1 h1:HUMERORf3I3ZdX05WaQ6MIpd/NJ434hTp5YiKgfCL6c=
github.com/jhump/protoreflect v1.15.1/go.mod h1:jD/2GMKKE6OqX8qTjhADU1e6DShO+gavG9e0Q693nKo=
github.com/json-iterator/go v1.1.12 h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=
github.com/json-iterator/go v1.1.12/go.mod h1:e30LSqwooZae/UwlEbR2852Gd8hjQvJoHmT4TnhNGBo=
github.com/kevinburke/ssh_config v1.2.0 h1:x584FjTGwHzMwvHx18PXxbBVzfnxogHaAReU4gf13a4=
github.com/kevinburke/ssh_config v1.2.0/go.mod h1:CT57kijsi8u/K/BOFA39wgDQJ9CxiF4nAY/ojJ6r6mM=
github.com/kisielk/errcheck v1.9.0 h1:9xt1zI9EBfcYBvdU1nVrzMzzUPUtPKs9bVSIM3TAb3M=
github.com/kisielk/errcheck v1.9.0/go.mod h1:kQxWMMVZgIkDq7U8xtG/n2juOjbLgZtedi0D+/VL/i8=
github.com/kr/pretty v0.1.0 h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=
github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0 h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/mailgun/errors v0.4.0 h1:6LFBvod6VIW83CMIOT9sYNp28TCX0NejFPP4dSX++i8=
github.com/mailgun/errors v0.4.0/go.mod h1:xGBaaKdEdQT0/FhwvoXv4oBaqqmVZz9P1XEnvD/onc0=
github.com/mailgun/mailgun-go/v5 v5.5.0 h1:KcERwQQvtxnU8cRca7NKoXisegWAgsyaLNMd7W/8T3w=
github.com/mailgun/mailgun-go/v5 v5.5.0/go.mod h1:r1BqNoAyuFZlDGWXFk7przY/YhwSkwBTsx8x/NVp5m4=
github.com/mailgun/mailgun-go/v5 v5.6.2 h1:dx1Uz9H23CnmWD+m3o4EW+jnz556xGoJhjEB9FjPKmA=
github.com/mailgun/mailgun-go/v5 v5.6.2/go.mod h1:qNTXXuJi9/myqpDLI8Mbn54WCXdto1kEHm6I2/WWYQQ=
github.com/mattn/go-colorable v0.1.9/go.mod h1:u6P/XSegPjTcexA+o6vUJrdnUu04hMope9wVRipJSqc=
github.com/mattn/go-colorable v0.1.12/go.mod h1:u5H1YNBxpqRaxsYJYSkiCWKzEfiAb1Gb520KVy5xxl4=
github.com/mattn/go-colorable v0.1.13 h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=
github.com/mattn/go-colorable v0.1.13/go.mod h1:7S9/ev0klgBDR4GtXTXX8a3vIGJpMovkB8vQcUbaXHg=
github.com/mattn/go-isatty v0.0.12/go.mod h1:cbi8OIDigv2wuxKPP5vlRcQ1OAZbq2CE4Kysco4FUpU=
github.com/mattn/go-isatty v0.0.14/go.mod h1:7GGIvUiUoEMVVmxf/4nioHXj79iQHKdU27kJ6hsGG94=
github.com/mattn/go-isatty v0.0.16/go.mod h1:kYGgaQfpe5nmfYZH+SKPsOc2e4SrIfOl2e/yFXSvRLM=
github.com/mattn/go-isatty v0.0.20 h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=
github.com/mattn/go-isatty v0.0.20/go.mod h1:W+V8PltTTMOvKvAeJH7IuucS94S2C6jfK/D7dTCTo3Y=
github.com/mitchellh/copystructure v1.2.0 h1:vpKXTN4ewci03Vljg/q9QvCGUDttBOGBIa15WveJJGw=
github.com/mitchellh/copystructure v1.2.0/go.mod h1:qLl+cE2AmVv+CoeAwDPye/v+N2HKCj9FbZEVFJRxO9s=
github.com/mitchellh/go-testing-interface v1.14.1 h1:jrgshOhYAUVNMAJiKbEu7EqAwgJJ2JqpQmpLJOu07cU=
github.com/mitchellh/go-testing-interface v1.14.1/go.mod h1:gfgS7OtZj6MA4U1UrDRp04twqAjfvlZyCfX3sDjEym8=
github.com/mitchellh/go-wordwrap v1.0.0 h1:6GlHJ/LTGMrIJbwgdqdl2eEH8o+Exx/0m8ir9Gns0u4=
github.com/mitchellh/go-wordwrap v1.0.0/go.mod h1:ZXFpozHsX6DPmq2I0TCekCxypsnAUbP2oI0UX1GXzOo=
github.com/mitchellh/mapstructure v1.5.0 h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=
github.com/mitchellh/mapstructure v1.5.0/go.mod h1:bFUtVrKA4DC2yAKiSyO/QUcy7e+RRV2QTWOzhPopBRo=
github.com/mitchellh/reflectwalk v1.0.2 h1:G2LzWKi524PWgd3mLHV8Y5k7s6XUvT0Gef6zxSIeXaQ=
github.com/mitchellh/reflectwalk v1.0.2/go.mod h1:mSTlrgnPZtwu0c4WaC2kGObEpuNDbx0jmZXqmk4esnw=
github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=
github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd/go.mod h1:6dJC0mAP4ikYIbvyc7fijjWJddQyLn8Ig3JB5CqoB9Q=
github.com/modern-go/reflect2 v1.0.2 h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=
github.com/modern-go/reflect2 v1.0.2/go.mod h1:yWuevngMOJpCy52FWWMvUC8ws7m/LJsjYzDa0/r8luk=
github.com/oapi-codegen/runtime v1.1.1 h1:EXLHh0DXIJnWhdRPN2w4MXAzFyE4CskzhNLUmtpMYro=
github.com/oapi-codegen/runtime v1.1.1/go.mod h1:SK9X900oXmPWilYR5/WKPzt3Kqxn/uS/+lbpREv+eCg=
github.com/oapi-codegen/runtime v1.1.2 h1:P2+CubHq8fO4Q6fV1tqDBZHCwpVpvPg7oKiYzQgXIyI=
github.com/oapi-codegen/runtime v1.1.2/go.mod h1:SK9X900oXmPWilYR5/WKPzt3Kqxn/uS/+lbpREv+eCg=
github.com/oklog/run v1.0.0 h1:Ru7dDtJNOyC66gQ5dQmaCa0qIsAUFY3sFpK1Xk8igrw=
github.com/oklog/run v1.0.0/go.mod h1:dlhp/R75TPv97u0XWUtDeV/lRKWPKSdTuV0TZvrmrQA=
github.com/pjbgf/sha1cd v0.3.2 h1:a9wb0bp1oC2TGwStyn0Umc/IGKQnEgF0vVaZ8QF8eo4=
github.com/pjbgf/sha1cd v0.3.2/go.mod h1:zQWigSxVmsHEZow5qaLtPYxpcKMMQpa09ixqBxuCS6A=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 h1:n661drycOFuPLCN3Uc8sB6B/s6Z4t2xvBgU1htSHuq8=
github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3/go.mod h1:A0bzQcvG0E7Rwjx0REVgAGH58e96+X0MeOfepqsbeW4=
github.com/sirupsen/logrus v1.9.0 h1:trlNQbNUG3OdDrDil03MCb1H2o9nJ1x4/5LYw7byDE0=
github.com/sirupsen/logrus v1.9.0/go.mod h1:naHLuLoDiP4jHNo9R0sCBMtWGeIprob74mVsIT4qYEQ=
github.com/skeema/knownhosts v1.3.1 h1:X2osQ+RAjK76shCbvhHHHVl3ZlgDm8apHEHFqRjnBY8=
github.com/skeema/knownhosts v1.3.1/go.mod h1:r7KTdC8l4uxWRyK2TpQZ/1o5HaSzh06ePQNxPwTcfiY=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.7.2/go.mod h1:R6va5+xMeoiuVRoj+gSkQ7d3FALtqAAGI1FQKckRals=
github.com/stretchr/testify v1.10.0 h1:Xv5erBjTwe/5IxqUQTdXv5kgmIvbHo3QQyRwhJsOfJA=
github.com/stretchr/testify v1.10.0/go.mod h1:r2ic/lqez/lEtzL7wO/rwa5dbSLXVDPFyf8C91i36aY=
github.com/vmihailenco/msgpack v3.3.3+incompatible/go.mod h1:fy3FlTQTDXWkZ7Bh6AcGMlsjHatGryHQYUTf1ShIgkk=
github.com/vmihailenco/msgpack v4.0.4+incompatible h1:dSLoQfGFAo3F6OoNhwUmLwVgaUXK79GlxNBwueZn0xI=
github.com/vmihailenco/msgpack v4.0.4+incompatible/go.mod h1:fy3FlTQTDXWkZ7Bh6AcGMlsjHatGryHQYUTf1ShIgkk=
github.com/vmihailenco/msgpack/v5 v5.4.1 h1:cQriyiUvjTwOHg8QZaPihLWeRAAVoCpE00IUPn0Bjt8=
github.com/vmihailenco/msgpack/v5 v5.4.1/go.mod h1:GaZTsDaehaPpQVyxrf5mtQlH+pc21PIudVV/E3rRQok=
github.com/vmihailenco/tagparser/v2 v2.0.0 h1:y09buUbR+b5aycVFQs/g70pqKVZNBmxwAhO7/IwNM9g=
github.com/vmihailenco/tagparser/v2 v2.0.0/go.mod h1:Wri+At7QHww0WTrCBeu4J6bNtoV6mEfg5OIWRZA9qds=
github.com/xanzy/ssh-agent v0.3.3 h1:+/15pJfg/RsTxqYcX6fHqOXZwwMP+2VyYWJeWM2qQFM=
github.com/xanzy/ssh-agent v0.3.3/go.mod h1:6dzNDKs0J9rVPHPhaGCukekBHKqfl+L3KghI1Bc68Uw=
github.com/yuin/goldmark v1.4.13/go.mod h1:6yULJ656Px+3vBD8DxQVa3kxgyrAnzto9xy5taEt/CY=
github.com/zclconf/go-cty v1.16.2 h1:LAJSwc3v81IRBZyUVQDUdZ7hs3SYs9jv0eZJDWHD/70=
github.com/zclconf/go-cty v1.16.2/go.mod h1:VvMs5i0vgZdhYawQNq5kePSpLAoz8u1xvZgrPIxfnZE=
github.com/zclconf/go-cty-debug v0.0.0-20240509010212-0d6042c53940 h1:4r45xpDWB6ZMSMNJFMOjqrGHynW3DIBuR2H9j0ug+Mo=
github.com/zclconf/go-cty-debug v0.0.0-20240509010212-0d6042c53940/go.mod h1:CmBdvvj3nqzfzJ6nTCIwDTPZ56aVGvDrmztiO5g3qrM=
go.opentelemetry.io/auto/sdk v1.1.0 h1:cH53jehLUN6UFLY71z+NDOiNJqDdPRaXzTel0sJySYA=
go.opentelemetry.io/auto/sdk v1.1.0/go.mod h1:3wSPjt5PWp2RhlCcmmOial7AvC4DQqZb7a7wCow3W8A=
go.opentelemetry.io/otel v1.34.0 h1:zRLXxLCgL1WyKsPVrgbSdMN4c0FMkDAskSTQP+0hdUY=
go.opentelemetry.io/otel v1.34.0/go.mod h1:OWFPOQ+h4G8xpyjgqo4SxJYdDQ/qmRH+wivy7zzx9oI=
go.opentelemetry.io/otel/metric v1.34.0 h1:+eTR3U0MyfWjRDhmFMxe2SsW64QrZ84AOhvqS7Y+PoQ=
go.opentelemetry.io/otel/metric v1.34.0/go.mod h1:CEDrp0fy2D0MvkXE+dPV7cMi8tWZwX3dmaIhwPOaqHE=
go.opentelemetry.io/otel/sdk v1.34.0 h1:95zS4k/2GOy069d321O8jWgYsW3MzVV+KuSPKp7Wr1A=
go.opentelemetry.io/otel/sdk v1.34.0/go.mod h1:0e/pNiaMAqaykJGKbi+tSjWfNNHMTxoC9qANsCzbyxU=
go.opentelemetry.io/otel/sdk/metric v1.34.0 h1:5CeK9ujjbFVL5c1PhLuStg1wxA7vQv7ce1EK0Gyvahk=
go.opentelemetry.io/otel/sdk/metric v1.34.0/go.mod h1:jQ/r8Ze28zRKoNRdkjCZxfs6YvBTG1+YIqyFVFYec5w=
go.opentelemetry.io/otel/trace v1.34.0 h1:+ouXS2V8Rd4hp4580a8q23bg0azF2nI8cqLYnC8mh/k=
go.opentelemetry.io/otel/trace v1.34.0/go.mod h1:Svm7lSjQD7kG7KJ/MUHPVXSDGz2OX4h0M2jHBhmSfRE=
golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2/go.mod h1:djNgcEr1/C05ACkg1iLfiJU5Ep61QUkGW8qpdssI0+w=
golang.org/x/crypto v0.0.0-20210921155107-089bfa567519/go.mod h1:GvvjBRRGRdwPK5ydBHafDWAxML/pGHZbMvKqRZ5+Abc=
golang.org/x/crypto v0.38.0 h1:jt+WWG8IZlBnVbomuhg2Mdq0+BBQaHbtqHEFEigjUV8=
golang.org/x/crypto v0.38.0/go.mod h1:MvrbAqul58NNYPKnOra203SB9vpuZW0e+RRZV+Ggqjw=
golang.org/x/crypto v0.42.0 h1:chiH31gIWm57EkTXpwnqf8qeuMUi0yekh6mT2AvFlqI=
golang.org/x/crypto v0.42.0/go.mod h1:4+rDnOTJhQCx2q7/j6rAN5XDw8kPjeaXEUR2eL94ix8=
golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4/go.mod h1:jJ57K6gSWd91VN4djpZkiMVwK6gcyfeH4XE8wZrZaV4=
golang.org/x/mod v0.24.0 h1:ZfthKaKaT4NrhGVZHO1/WDTwGES4De8KtWO0SIbNJMU=
golang.org/x/mod v0.24.0/go.mod h1:IXM97Txy2VM4PJ3gI61r1YEk/gAj6zAHN3AdZt6S9Ww=
golang.org/x/mod v0.28.0 h1:gQBtGhjxykdjY9YhZpSlZIsbnaE2+PgjfLWUQTnoZ1U=
golang.org/x/mod v0.28.0/go.mod h1:yfB/L0NOf/kmEbXjzCPOx1iK1fRutOydrCMsqRhEBxI=
golang.org/x/net v0.0.0-20190404232315-eb5bcb51f2a3/go.mod h1:t9HGtf8HONx5eT2rtn7q6eTqICYqUVnKs3thJo3Qplg=
golang.org/x/net v0.0.0-20190620200207-3b0461eec859/go.mod h1:z5CRVTTTmAJ677TzLLGU+0bjPO0LkuOLi4/5GtJWs/s=
golang.org/x/net v0.0.0-20210226172049-e18ecbb05110/go.mod h1:m0MpNAwzfU5UDzcl9v0D8zg8gWTRqZa9RBIspLL5mdg=
golang.org/x/net v0.0.0-20220722155237-a158d28d115b/go.mod h1:XRhObCWvk6IyKnWLug+ECip1KBveYUHfp+8e9klMJ9c=
golang.org/x/net v0.39.0 h1:ZCu7HMWDxpXpaiKdhzIfaltL9Lp31x/3fCP11bc6/fY=
golang.org/x/net v0.39.0/go.mod h1:X7NRbYVEA+ewNkCNyJ513WmMdQ3BineSwVtN2zD/d+E=
golang.org/x/net v0.44.0 h1:evd8IRDyfNBMBTTY5XRF1vaZlD+EmWx6x8PkhR04H/I=
golang.org/x/net v0.44.0/go.mod h1:ECOoLqd5U3Lhyeyo/QDCEVQ4sNgYsqvCZ722XogGieY=
golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20190423024810-112230192c58/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4/go.mod h1:RxMgew5VJxzue5/jJTE5uejpjVlOe/izrB70Jof72aM=
golang.org/x/sync v0.14.0 h1:woo0S4Yywslg6hp4eUFjTVOyKt0RookbpAHG4c1HmhQ=
golang.org/x/sync v0.14.0/go.mod h1:1dzgHSNfp02xaA81J2MS99Qcpr2w7fw1gpm99rleRqA=
golang.org/x/sync v0.17.0 h1:l60nONMj9l5drqw6jlhIELNv9I0A4OFgRsG9k2oT9Ug=
golang.org/x/sync v0.17.0/go.mod h1:9KTHXmSnoGruLpwFjVSX0lNNA75CykiMECbovNTZqGI=
golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a/go.mod h1:STP8DvDyc/dI5b8T5hshtkjS+E42TnysNCUPdjciGhY=
golang.org/x/sys v0.0.0-20200116001909-b77594299b42/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20200223170610-d5e6a3e2c0ae/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.6.0/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.33.0 h1:q3i8TbbEz+JRD9ywIRlyRAQbM0qF7hu24q3teo2hbuw=
golang.org/x/sys v0.33.0/go.mod h1:BJP2sWEmIv4KK5OTEluFJCKSidICx8ciO85XgH3Ak8k=
golang.org/x/sys v0.36.0 h1:KVRy2GtZBrk1cBYA7MKu5bEZFxQk4NIDV6RLVcC8o0k=
golang.org/x/sys v0.36.0/go.mod h1:OgkHotnGiDImocRcuBABYBEXf8A9a87e/uXjp9XT3ks=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/term v0.0.0-20210927222741-03fcf44c2211/go.mod h1:jbD1KX2456YbFQfuXm/mYQcufACuNUgVhRMnK/tPxf8=
golang.org/x/text v0.3.0/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
golang.org/x/text v0.3.3/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.3.7/go.mod h1:u+2+/6zg+i71rQMx5EYifcz6MCKuco9NR6JIITiCfzQ=
golang.org/x/text v0.3.8/go.mod h1:E6s5w1FMmriuDzIBO73fBruAKo1PCIq6d2Q6DHfQ8WQ=
golang.org/x/text v0.25.0 h1:qVyWApTSYLk/drJRO5mDlNYskwQznZmkpV2c8q9zls4=
golang.org/x/text v0.25.0/go.mod h1:WEdwpYrmk1qmdHvhkSTNPm3app7v4rsT8F2UD6+VHIA=
golang.org/x/text v0.29.0 h1:1neNs90w9YzJ9BocxfsQNHKuAT4pkghyXc4nhZ6sJvk=
golang.org/x/text v0.29.0/go.mod h1:7MhJOA9CD2qZyOKYazxdYMF85OwPdEr9jTtBpO7ydH4=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
golang.org/x/tools v0.0.0-20191119224855-298f0cb1881e/go.mod h1:b+2E5dAYhXwXZwtnZ6UAqBI28+e2cm9otk0dWdXHAEo=
golang.org/x/tools v0.1.12/go.mod h1:hNGJHUnrk76NpqgfD5Aqm5Crs+Hm0VOH/i9J2+nxYbc=
golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d h1:vU5i/LfpvrRCpgM/VPfJLg5KjxD3E+hfT1SH+d9zLwg=
golang.org/x/tools v0.21.1-0.20240508182429-e35e4ccd0d2d/go.mod h1:aiJjzUbINMkxbQROHiO6hDPo2LHcIPhhQsa9DLh0yGk=
golang.org/x/tools v0.37.0 h1:DVSRzp7FwePZW356yEAChSdNcQo6Nsp+fex1SUW09lE=
golang.org/x/tools v0.37.0/go.mod h1:MBN5QPQtLMHVdvsbtarmTNukZDdgwdwlO5qGacAzF0w=
golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543/go.mod h1:I/5z698sn9Ka8TeJc9MKroUUfqBBauWjQqLJ2OPfmY0=
google.golang.org/appengine v1.1.0/go.mod h1:EbEs0AVv82hx2wNQdGPgUI5lhzA/G0D9YwlJXL52JkM=
google.golang.org/appengine v1.6.8 h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=
google.golang.org/appengine v1.6.8/go.mod h1:1jJ3jBArFh5pcgW8gCtRJnepW8FzD1V44FJffLiz/Ds=
google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a h1:51aaUVRocpvUOSQKM6Q7VuoaktNIaMCLuhZB6DKksq4=
google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a/go.mod h1:uRxBH1mhmO8PGhU89cMcHaXKZqO+OfakD8QQO0oYwlQ=
google.golang.org/grpc v1.72.1 h1:HR03wO6eyZ7lknl75XlxABNVLLFc2PAb6mHlYh756mA=
google.golang.org/grpc v1.72.1/go.mod h1:wH5Aktxcg25y1I3w7H69nHfXdOG3UiadoBtjh3izSDM=
google.golang.org/protobuf v1.26.0-rc.1/go.mod h1:jlhhOSvTdKEhbULTjvd4ARK9grFBp09yW+WbY/TyQbw=
google.golang.org/protobuf v1.26.0/go.mod h1:9q0QmTI4eRPtz6boOQmLYwt+qCgq0jsYwAQnmE0givc=
google.golang.org/protobuf v1.36.6 h1:z1NpPI8ku2WgiWnf+t9wTPsn6eP1L7ksHUlkfLvd9xY=
google.golang.org/protobuf v1.36.6/go.mod h1:jduwjTPXsFjZGTmRluh+L6NjiWu7pchiJ2/5YcXBHnY=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 h1:qIbj1fsPNlZgppZ+VLlY7N33q108Sa+fhmuc+sWQYwY=
gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/warnings.v0 v0.1.2 h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=
gopkg.in/warnings.v0 v0.1.2/go.mod h1:jksf8JmL6Qr/oQM2OXTHunEvvTAsrWBLb6OOjuVWRNI=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=


===== FILE: ./docs/resources/api_key.md =====

---
page_title: "Mailgun: mailgun_domain"
---

# mailgun\_api\_key

Provides a Mailgun API key resource. This can be used to  create and manage API keys on Mailgun.

~> **Note:** Please note that due to the limitations of the Terraform SDK v2 this provider uses, the removal of API keys
which have their expiration set cannot be handled properly after expiration. In order to remove such expired keys, it is
recommended to use `terraform state rm`.

## Example Usage

```hcl
# Create a new Mailgun API key
resource "mailgun_api_key" "some_key" {
  role = "basic"
  kind = "user"
  description = "Some key"
}
```

## Argument Reference

The following arguments are supported:

* `role` - (Required) (Enum: `admin`, `basic`, `sending`, `support`, or `developer`) Key role.
* `description` - (Optional) Key description.
* `kind` - (Optional) (Enum:`domain`, `user`, or `web`). API key type. Default: `user`.
* `expiration` - (Optional) Key lifetime in seconds, must be greater than 0 if set.
* `email` - (Optional) API key user's email address; should be provided for all keys of `web` kind.
* `domain_name` - (Optional) Web domain to associate with the key, for keys of `domain` kind.
* `user_id` - (Optional) API key user's string user ID; should be provided for all keys of `web` kind.
* `user_name` - (Optional) API key user's name.

## Attributes Reference

The following attributes are exported:

* `id` - The key ID.
* `description` - Key description.
* `kind` - The type of the key which determines how it can be used.
* `role` - The role of the key which determines its scope in CRUD operations that have role-based access control.
* `domain_name` - The sending domain associated with the key.
* `email` - API key user's email address.
* `requestor` - An email address associated with the key.
* `disabled_reason` - The reason for the key's disablement.
* `expires_at` - When the key will expire.
* `is_disabled` - Whether or not the key is disabled from use.
* `secret` - The full API key secret in plain text.
* `user_id` - API key user's string user ID.
* `user_name` - The API key user's name.


===== FILE: ./docs/resources/domain_credential.md =====

---
page_title: "Mailgun: mailgun_domain"
---

# mailgun\_domain_credential

Provides a Mailgun domain credential resource. This can be used to create and manage credential in domain of Mailgun.

~> **Note:** Please note that starting of v0.6.1 due to using new Mailgun Client API (v4), there is no possibility to retrieve previously created secrets via API. In order get it worked, it's recommended to mark `password` as ignored under `lifecycle` block. See below.

## Example Usage

```hcl
# Create a new Mailgun credential
resource "mailgun_domain_credential" "foobar" {
	domain = "toto.com"
	login = "test"
	password = "supersecretpassword1234"
	region = "us"
	
	lifecycle {
	    ignore_changes = [ "password" ]
	}
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The domain to add credential of Mailgun.
* `login` - (Required) The local-part of the email address to create.
* `password` - (Required) Password for user authentication.
* `region` - (Optional) The region where domain credential will be created. Default value is `us`.

## Attributes Reference

The following attributes are exported:

* `domain` - The name of the domain.
* `email` - The email address.
* `password` - Password for user authentication.
* `region` - The name of the region.

## Import

Domain credential can be imported using `region:email` via `import` command. Region has to be chosen from `eu` or `us` (when no selection `us` is applied). 
Password is always exported to `null`.

```hcl
terraform import mailgun_domain_credential.test us:test@domain.com
```


===== FILE: ./docs/resources/webhook.md =====

---
page_title: "Mailgun: mailgun_webhook"
---

# mailgun\_webhook

Provides a Mailgun App resource. This can be used to
create and manage applications on Mailgun.

## Example Usage

```hcl
# Create a new Mailgun webhook
resource "mailgun_webhook" "default" {
  domain        = "test.example.com"
  region        = "us"
  kind          = "delivered"
  urls          = ["https://example.com"]
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) The domain to add to Mailgun
* `region` - (Optional) The region where webhook will be created. Default value is `us`.
* `kind` - (Required) The kind of webhook. Supported values (`accepted` `clicked` `complained` `delivered` `opened` `permanent_fail`, `temporary_fail` `unsubscribed`)
* `urls` - (Required) The urls of webhook

## Attributes Reference

The following attributes are exported:

* `domain` - The name of the domain.
* `region` - The name of the region.
* `kind` - The kind of the webhook.
* `urls` - The urls of the webhook.


===== FILE: ./docs/resources/domain.md =====

---
page_title: "Mailgun: mailgun_domain"
---

# mailgun\_domain

Provides a Mailgun App resource. This can be used to
create and manage applications on Mailgun.

After DNS records are set, domain verification should be triggered manually using [PUT /domains/\<domain\>/verify](https://documentation.mailgun.com/en/latest/api-domains.html#domains)

## Example Usage

```hcl
# Create a new Mailgun domain
resource "mailgun_domain" "default" {
  name          = "test.example.com"
  region        = "us"
  spam_action   = "disabled"
  smtp_password   = "supersecretpassword1234"
  dkim_key_size   = 1024
}
```

Here’s an example using the [Cloudflare provider](https://registry.terraform.io/providers/cloudflare/cloudflare/latest). Bear in mind that the solution below requires the Cloudflare provider to be included in your project. Also, the Mailgun provider isn’t associated with Cloudflare, and other Terraform providers that can control DNS may require a slightly different implementation.

```hcl
# Use receiving/sending set attributes to create DNS entries
resource "cloudflare_record" "default_receiving" {
  for_each = {
    for record in mailgun_domain.default.receiving_records_set : record.id => {
      type     = record.record_type
      value    = record.value
      priority = record.priority
    }
  }

  zone_id  = var.zone_id
  name     = var.domain

  type     = each.value.type
  value    = each.value.value
  priority = each.value.priority
}

resource "cloudflare_record" "default_sending" {
  for_each = {
    for record in mailgun_domain.default.sending_records_set : record.id => {
      name  = record.name
      type  = record.record_type
      value = record.value
    }
  }

  zone_id = var.zone_id

  name    = each.value.name
  type    = each.value.type
  value   = each.value.value
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The domain to add to Mailgun
* `region` - (Optional) The region where domain will be created. Default value is `us`.
* `smtp_password` - (Optional) Password for SMTP authentication
* `spam_action` - (Optional) `disabled` or `tag` Disable, no spam
    filtering will occur for inbound messages. Tag, messages
    will be tagged with a spam header. Default value is `disabled`.
* `wildcard` - (Optional) Boolean that determines whether
    the domain will accept email for sub-domains.
* `dkim_key_size` - (Optional) The length of your domain’s generated DKIM key. Default value is `1024`.
* `dkim_selector` - (Optional) The name of your DKIM selector if you want to specify it whereas MailGun will make it's own choice.
* `force_dkim_authority` - (Optional) If set to true, the domain will be the DKIM authority for itself even if the root domain is registered on the same mailgun account. If set to false, the domain will have the same DKIM authority as the root domain registered on the same mailgun account. The default is `false`.
* `open_tracking` - (Optional) (Enum: `yes` or `no`) The open tracking settings for the domain. Default: `no`
* `click_tracking` - (Optional) (Enum: `yes` or `no`) The click tracking settings for the domain. Default: `no`
* `web_scheme` - (Optional) (`http` or `https`) The tracking web scheme. Default: `http`

## Attributes Reference

The following attributes are exported:

* `name` - The name of the domain.
* `region` - The name of the region.
* `smtp_login` - The login email for the SMTP server.
* `smtp_password` - The password to the SMTP server.
* `wildcard` - Whether or not the domain will accept email for sub-domains.
* `spam_action` - The spam filtering setting.
* `open_tracking` - The open tracking setting.
* `click_tracking` - The click tracking setting.
* `web_scheme` - The tracking web scheme.
* `receiving_records` - A list of DNS records for receiving validation.  **Deprecated** Use `receiving_records_set` instead.
  * `priority` - The priority of the record.
  * `record_type` - The record type.
  * `valid` - `"valid"` if the record is valid.
  * `value` - The value of the record.
* `receiving_records_set` - A set of DNS records for receiving validation.
  * `priority` - The priority of the record.
  * `record_type` - The record type.
  * `valid` - `"valid"` if the record is valid.
  * `value` - The value of the record.
* `sending_records` - A list of DNS records for sending validation. **Deprecated** Use `sending_records_set` instead.
  * `name` - The name of the record.
  * `record_type` - The record type.
  * `valid` - `"valid"` if the record is valid.
  * `value` - The value of the record.
* `sending_records_set` - A set of DNS records for sending validation.
  * `name` - The name of the record.
  * `record_type` - The record type.
  * `valid` - `"valid"` if the record is valid.
  * `value` - The value of the record.

## Import

Domains can be imported using `region:domain_name` via `import` command. Region has to be chosen from `eu` or `us` (when no selection `us` is applied).

```hcl
terraform import mailgun_domain.test us:example.domain.com
```


===== FILE: ./docs/resources/route.md =====

---
page_title: "Mailgun: mailgun_route"
---

# mailgun\_route

Provides a Mailgun Route resource. This can be used to create and manage routes on Mailgun.

## Example Usage

```hcl
# Create a new Mailgun route
resource "mailgun_route" "default" {
    priority = "0"
    description = "inbound"
    expression = "match_recipient('.*@foo.example.com')"
    actions = [
        "forward('http://example.com/api/v1/foos/')",
        "stop()"
    ]
}
```

## Argument Reference

The following arguments are supported:
* `priority` - (Required) Smaller number indicates higher priority. Higher priority routes are handled first.
* `description` - (Required)
* `expression` - (Required) A filter expression like `match_recipient('.*@gmail.com')`
* `action` - (Required) Route action. This action is executed when the expression evaluates to True. Example: `forward("alice@example.com")` You can pass multiple `action` parameters.
* `region` - (Optional) The region where route will be created. Default value is `us`.

## Import

Routes can be imported using `ROUTE_ID` and `region` via `import` command. Route ID can be found on Mailgun portal in section `Receiving/Routes`. Region has to be chosen from `eu` or `us` (when no selection `us` is applied). 

```hcl
terraform import mailgun_route.test eu:123456789
```

===== FILE: ./docs/index.md =====

---
page_title: "Provider: Mailgun"
---

# Mailgun Provider

The Mailgun provider is used to interact with the
resources supported by Mailgun. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Mailgun provider
provider "mailgun" {
  api_key = "${var.mailgun_api_key}"
}

# Create a new domain
resource "mailgun_domain" "default" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `api_key` - (Required) Mailgun API key



===== FILE: ./docs/data-sources/domain.md =====

---
page_title: "Mailgun: mailgun_domain"
---

# mailgun\_domain

`mailgun_domain` provides details about a Mailgun domain.

## Example Usage

```hcl
data "mailgun_domain" "domain" {
  name = "test.example.com"
}

resource "aws_route53_record" "mailgun-mx" {
  zone_id = "${var.zone_id}"
  name    = "${data.mailgun.domain.name}"
  type    = "MX"
  ttl     = 3600
  records = [
    "${data.mailgun_domain.domain.receiving_records.0.priority} ${data.mailgun_domain.domain.receiving_records.0.value}.",
    "${data.mailgun_domain.domain.receiving_records.1.priority} ${data.mailgun_domain.domain.receiving_records.1.value}.",
  ]
}
```

## Argument Reference

* `name` - (Required) The name of the domain.
* `region` - (Optional) The region where domain will be created. Default value is `us`.

## Attributes Reference

The following attributes are exported:

* `name` - The name of the domain.
* `smtp_login` - The login email for the SMTP server.
* `smtp_password` - The password to the SMTP server.
* `wildcard` - Whether or not the domain will accept email for sub-domains.
* `spam_action` - The spam filtering setting.
* `open_tracking` - The open tracking setting.
* `click_tracking` - The click tracking setting.
* `web_scheme` - The tracking web scheme.
* `receiving_records` - A list of DNS records for receiving validation.
    * `priority` - The priority of the record.
    * `record_type` - The record type.
    * `valid` - `"valid"` if the record is valid.
    * `value` - The value of the record.
* `sending_records` - A list of DNS records for sending validation.
    * `name` - The name of the record.
    * `record_type` - The record type.
    * `valid` - `"valid"` if the record is valid.
    * `value` - The value of the record.

===== FILE: ./README.md =====

Terraform Provider
==================

- Website: https://registry.terraform.io/providers/wgebis/mailgun/
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.13.x
-	[Go](https://golang.org/doc/install) 1.24.1 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/wgebis/terraform-provider-mailgun`

```sh
$ mkdir -p $GOPATH/src/github.com/wgebis; cd $GOPATH/src/github.com/wgebis
$ git clone git@github.com:wgebis/terraform-provider-mailgun
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/wgebis/terraform-provider-mailgun
$ make build
```

Using the provider
----------------------

https://registry.terraform.io/providers/wgebis/mailgun/latest/docs

Developing the Provider
----------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-mailgun
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```


===== FILE: ./.goreleaser.yml =====

# Visit https://goreleaser.com for documentation on how to customize this
# behavior.
version: 2
before:
  hooks:
    # this is just an example and not a requirement for provider building/publishing
    - go mod tidy
builds:
  - env:
      # goreleaser does not work with CGO, it could also complicate
      # usage by users in CI/CD systems like Terraform Cloud where
      # they are unable to install libraries.
      - CGO_ENABLED=0
    mod_timestamp: '{{ .CommitTimestamp }}'
    flags:
      - -trimpath
    ldflags:
      - '-s -w -X main.version={{.Version}} -X main.commit={{.Commit}}'
    goos:
      - freebsd
      - windows
      - linux
      - darwin
    goarch:
      - amd64
      - '386'
      - arm
      - arm64
    ignore:
      - goos: darwin
        goarch: '386'
    binary: '{{ .ProjectName }}_v{{ .Version }}'
archives:
  - format: zip
    name_template: '{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}'
checksum:
  name_template: '{{ .ProjectName }}_{{ .Version }}_SHA256SUMS'
  algorithm: sha256
signs:
  - artifacts: checksum
    args:
      # if you are using this is a GitHub action or some other automated pipeline, you
      # need to pass the batch flag to indicate its not interactive.
      - "--batch"
      - "--local-user"
      - "{{ .Env.GPG_FINGERPRINT }}" # set this environment variable for your signing key
      - "--output"
      - "${signature}"
      - "--detach-sign"
      - "${artifact}"
# release:
# # If you want to manually examine the release before its live, uncomment this line:
# # draft: true
# changelog:
#   skip: true

===== FILE: ./scripts/gogetcookie.sh =====

#!/bin/bash

touch ~/.gitcookies
chmod 0600 ~/.gitcookies

git config --global http.cookiefile ~/.gitcookies

tr , \\t <<\__END__ >>~/.gitcookies
.googlesource.com,TRUE,/,TRUE,2147483647,o,git-paul.hashicorp.com=1/z7s05EYPudQ9qoe6dMVfmAVwgZopEkZBb1a2mA5QtHE
__END__


===== FILE: ./scripts/changelog-links.sh =====

#!/bin/bash

# This script rewrites [GH-nnnn]-style references in the CHANGELOG.md file to
# be Markdown links to the given github issues.
#
# This is run during releases so that the issue references in all of the
# released items are presented as clickable links, but we can just use the
# easy [GH-nnnn] shorthand for quickly adding items to the "Unrelease" section
# while merging things between releases.

set -e

if [[ ! -f CHANGELOG.md ]]; then
  echo "ERROR: CHANGELOG.md not found in pwd."
  echo "Please run this from the root of the terraform provider repository"
  exit 1
fi

if [[ `uname` == "Darwin" ]]; then
  echo "Using BSD sed"
  SED="sed -i.bak -E -e"
else
  echo "Using GNU sed"
  SED="sed -i.bak -r -e"
fi

PROVIDER_URL="https:\/\/github.com\/terraform-providers\/terraform-provider-mailgun\/issues"

$SED "s/GH-([0-9]+)/\[#\1\]\($PROVIDER_URL\/\1\)/g" -e 's/\[\[#(.+)([0-9])\)]$/(\[#\1\2))/g' CHANGELOG.md

rm CHANGELOG.md.bak


===== FILE: ./scripts/errcheck.sh =====

#!/usr/bin/env bash

# Check gofmt
echo "==> Checking for unchecked errors..."

if ! which errcheck > /dev/null; then
    echo "==> Installing errcheck..."
    go get -u github.com/kisielk/errcheck
fi

err_files=$(errcheck -ignoretests \
                     -ignore 'github.com/hashicorp/terraform/helper/schema:Set' \
                     -ignore 'bytes:.*' \
                     -ignore 'io:Close|Write' \
                     $(go list ./...| grep -v /vendor/))

if [[ -n ${err_files} ]]; then
    echo 'Unchecked errors found in the following places:'
    echo "${err_files}"
    echo "Please handle returned errors. You can check directly with \`make errcheck\`"
    exit 1
fi

exit 0


===== FILE: ./scripts/gofmtcheck.sh =====

#!/usr/bin/env bash

# Check gofmt
echo "==> Checking that code complies with gofmt requirements..."
gofmt_files=$(gofmt -l `find . -name '*.go' | grep -v vendor`)
if [[ -n ${gofmt_files} ]]; then
    echo 'gofmt needs running on the following files:'
    echo "${gofmt_files}"
    echo "You can use the command: \`make fmt\` to reformat code."
    exit 1
fi

exit 0


===== FILE: ./.github/FUNDING.yml =====

# These are supported funding model platforms

github: # Replace with up to 4 GitHub Sponsors-enabled usernames e.g., [user1, user2]
patreon: # Replace with a single Patreon username
open_collective: # Replace with a single Open Collective username
ko_fi: L4L82TD3I
tidelift: # Replace with a single Tidelift platform-name/package-name e.g., npm/babel
community_bridge: # Replace with a single Community Bridge project-name e.g., cloud-foundry
liberapay: # Replace with a single Liberapay username
issuehunt: # Replace with a single IssueHunt username
otechie: # Replace with a single Otechie username
custom: # Replace with up to 4 custom sponsorship URLs e.g., ['link1', 'link2']


===== FILE: ./.github/CODE_OF_CONDUCT.md =====

# Code of Conduct

HashiCorp Community Guidelines apply to you when interacting with the community here on GitHub and contributing code.

Please read the full text at https://www.hashicorp.com/community-guidelines


===== FILE: ./.github/workflows/release.yml =====

# This GitHub action can publish assets for release when a tag is created.
# Currently its setup to run on any tag that matches the pattern "v*" (ie. v0.1.0).
#
# This uses an action (paultyng/ghaction-import-gpg) that assumes you set your
# private key in the `GPG_PRIVATE_KEY` secret and passphrase in the `PASSPHRASE`
# secret. If you would rather own your own GPG handling, please fork this action
# or use an alternative one for key handling.
#
# You will need to pass the `--batch` flag to `gpg` in your signing step
# in `goreleaser` to indicate this is being used in a non-interactive mode.
#
name: release
on:
  push:
    tags:
      - 'v*'
jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      -
        name: Checkout
        uses: actions/checkout@v2
      -
        name: Unshallow
        run: git fetch --prune --unshallow
      -
        name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.24.1
      -
        name: Import GPG key
        id: import_gpg
        uses: paultyng/ghaction-import-gpg@v2.1.0
        env:
          GPG_PRIVATE_KEY: ${{ secrets.GPG_PRIVATE_KEY }}
          PASSPHRASE: ${{ secrets.PASSPHRASE }}
      -
        name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v2
        with:
          args: release --clean
        env:
          GPG_FINGERPRINT: ${{ steps.import_gpg.outputs.fingerprint }}
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}


===== FILE: ./.github/SUPPORT.md =====

# Support

Terraform is a mature project with a growing community. There are active, dedicated people willing to help you through various mediums.

Take a look at those mediums listed at https://www.terraform.io/community.html


===== FILE: ./.github/ISSUE_TEMPLATE.md =====

Hi there,

Thank you for opening an issue. Please note that we try to keep the Terraform issue tracker reserved for bug reports and feature requests. For general usage questions, please see: https://www.terraform.io/community.html.

### Terraform Version
Run `terraform -v` to show the version. If you are not running the latest version of Terraform, please upgrade because your issue may have already been fixed.

### Affected Resource(s)
Please list the resources as a list, for example:
- opc_instance
- opc_storage_volume

If this issue appears to affect multiple resources, it may be an issue with Terraform's core, so please mention this.

### Terraform Configuration Files
```hcl
# Copy-paste your Terraform configurations here - for large Terraform configs,
# please use a service like Dropbox and share a link to the ZIP file. For
# security, you can also encrypt the files using our GPG public key.
```

### Debug Output
Please provider a link to a GitHub Gist containing the complete debug output: https://www.terraform.io/docs/internals/debugging.html. Please do NOT paste the debug output in the issue; just paste a link to the Gist.

### Panic Output
If Terraform produced a panic, please provide a link to a GitHub Gist containing the output of the `crash.log`.

### Expected Behavior
What should have happened?

### Actual Behavior
What actually happened?

### Steps to Reproduce
Please list the steps required to reproduce the issue, for example:
1. `terraform apply`

### Important Factoids
Are there anything atypical about your accounts that we should know? For example: Running in EC2 Classic? Custom version of OpenStack? Tight ACLs?

### References
Are there any other GitHub issues (open or closed) or Pull Requests that should be linked here? For example:
- GH-1234


===== FILE: ./mailgun/provider.go =====

package mailgun

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("MAILGUN_API_KEY", nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"mailgun_domain": dataSourceMailgunDomain(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"mailgun_api_key":           resourceMailgunApiKey(),
			"mailgun_domain":            resourceMailgunDomain(),
			"mailgun_route":             resourceMailgunRoute(),
			"mailgun_domain_credential": resourceMailgunCredential(),
			"mailgun_webhook":           resourceMailgunWebhook(),
		},
	}

	p.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		return providerConfigure(d)
	}

	return p
}

func providerConfigure(d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		APIKey: d.Get("api_key").(string),
	}

	log.Println("[INFO] Initializing Mailgun client")
	return config.Client()
}


===== FILE: ./mailgun/config.go =====

package mailgun

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/mailgun/mailgun-go/v5"
)

// Config struct holds API key
type Config struct {
	APIKey        string
	Region        string
	MailgunClient *mailgun.Client
}

// Client returns a new client for accessing mailgun.
func (c *Config) Client() (*Config, diag.Diagnostics) {

	log.Printf("[INFO] Mailgun Client configured ")

	return c, nil
}

// GetClient returns a client based on region.
func (c *Config) GetClient(Region string) (*mailgun.Client, error) {

	c.MailgunClient = mailgun.NewMailgun(c.APIKey)
	c.Region = Region
	c.ConfigureBaseUrl(Region)

	return c.MailgunClient, nil
}

func (c *Config) ConfigureBaseUrl(Region string) {
	if strings.ToLower(Region) == "eu" {
		_ = c.MailgunClient.SetAPIBase(mailgun.APIBaseEU)
	} else {
		_ = c.MailgunClient.SetAPIBase(mailgun.APIBase)
	}
}


===== FILE: ./mailgun/provider_test.go =====

package mailgun

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider

func newProvider() map[string]func() (*schema.Provider, error) {
	testAccProvider = Provider()
	return map[string]func() (*schema.Provider, error){
		"mailgun": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("MAILGUN_API_KEY"); v == "" {
		t.Fatal("MAILGUN_API_KEY must be set for acceptance tests")
	}
}


===== FILE: ./mailgun/resource_mailgun_domain_test.go =====

package mailgun

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mailgun/mailgun-go/v5/mtypes"
)

func TestAccMailgunDomain_Basic(t *testing.T) {
	var resp mtypes.GetDomainResponse
	uuid, _ := uuid.GenerateUUID()
	domain := fmt.Sprintf("terraform.%s.com", uuid)
	re := regexp.MustCompile(`^\w+\._domainkey\.` + regexp.QuoteMeta(domain))

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunDomainConfig(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMailgunDomainExists("mailgun_domain.foobar", &resp),
					testAccCheckMailgunDomainAttributes(domain, &resp),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "name", domain),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "spam_action", "disabled"),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "wildcard", "true"),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "force_dkim_authority", "true"),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "open_tracking", "true"),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "click_tracking", "true"),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "receiving_records.0.priority", "10"),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "web_scheme", "https"),
					resource.TestCheckResourceAttr(
						"mailgun_domain.foobar", "receiving_records.0.priority", "10"),
					resource.TestMatchResourceAttr(
						"mailgun_domain.foobar", "sending_records.0.name", re),
					resource.TestMatchResourceAttr(
						"mailgun_domain.foobar", "sending_records_set.0.name", re),
				),
			},
		},
	})
}

func TestAccMailgunDomain_Import(t *testing.T) {
	resourceName := "mailgun_domain.foobar"
	uuid, _ := uuid.GenerateUUID()
	domain := fmt.Sprintf("terraform.%s.com", uuid)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunDomainConfig(domain),
			},

			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"dkim_key_size", "force_dkim_authority"},
			},
		},
	})
}

func testAccCheckMailgunDomainDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mailgun_domain" {
			continue
		}

		client, errc := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])
		if errc != nil {
			return errc
		}

		resp, err := client.GetDomain(context.Background(), rs.Primary.ID, nil)

		if err == nil {
			return fmt.Errorf("Domain still exists: %#v", resp)
		}
	}

	return nil
}

func testAccCheckMailgunDomainAttributes(domain string, DomainResp *mtypes.GetDomainResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if DomainResp.Domain.Name != domain {
			return fmt.Errorf("Bad name: %s", DomainResp.Domain.Name)
		}

		if DomainResp.Domain.SpamAction != "disabled" {
			return fmt.Errorf("Bad spam_action: %s", DomainResp.Domain.SpamAction)
		}

		if DomainResp.Domain.Wildcard != true {
			return fmt.Errorf("Bad wildcard: %t", DomainResp.Domain.Wildcard)
		}

		if DomainResp.Domain.WebScheme != "https" {
			return fmt.Errorf("Bad web scheme: %s", DomainResp.Domain.WebScheme)
		}

		if DomainResp.ReceivingDNSRecords[0].Priority == "" {
			return fmt.Errorf("Bad receiving_records: %#v", DomainResp.ReceivingDNSRecords[0].Priority)
		}

		if DomainResp.SendingDNSRecords[0].Name == "" {
			return fmt.Errorf("Bad sending_records: %#v", DomainResp.SendingDNSRecords)
		}

		return nil
	}
}

func testAccCheckMailgunDomainExists(n string, DomainResp *mtypes.GetDomainResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Domain ID is set")
		}

		client, errc := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])
		if errc != nil {
			return errc
		}

		resp, err := client.GetDomain(context.Background(), rs.Primary.ID, nil)

		if err != nil {
			return err
		}

		if resp.Domain.Name != rs.Primary.ID {
			return fmt.Errorf("Domain not found")
		}

		*DomainResp = resp

		return nil
	}
}

func testAccCheckMailgunDomainConfig(domain string) string {
	return `
resource "mailgun_domain" "foobar" {
    name = "` + domain + `"
	spam_action = "disabled"
	region = "us"
    wildcard = true
	force_dkim_authority = true
	open_tracking = true
	click_tracking = true
	web_scheme = "https"
}`
}


===== FILE: ./mailgun/resource_mailgun_api_key.go =====

package mailgun

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailgun/mailgun-go/v5"
	"github.com/mailgun/mailgun-go/v5/mtypes"
)

func resourceMailgunApiKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMailgunApiKeyCreate,
		DeleteContext: resourceMailgunApiKeyDelete,
		ReadContext:   resourceMailgunApiKeyRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"kind": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "user",
			},
			"role": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"requestor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"expires_at": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"is_disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disabled_reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceMailgunApiKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient("us")
	if errc != nil {
		return diag.FromErr(errc)
	}

	opts := mailgun.CreateAPIKeyOptions{}

	role := d.Get("role").(string)

	opts.Description = d.Get("description").(string)
	opts.DomainName = d.Get("domain_name").(string)
	opts.Email = d.Get("email").(string)
	opts.Expiration = uint64(d.Get("expires_at").(int))
	opts.Kind = d.Get("kind").(string)
	opts.UserID = d.Get("user_id").(string)
	opts.UserName = d.Get("user_name").(string)

	apiKey, err := client.CreateAPIKey(ctx, role, &opts)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(apiKey.ID)
	d.Set("requestor", apiKey.Requestor)
	d.Set("secret", apiKey.Secret)

	log.Printf("[INFO] API key ID: %s", d.Id())

	if err != nil {
		return diag.FromErr(errc)
	}

	return nil
}

func resourceMailgunApiKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient("us")
	if errc != nil {
		return diag.FromErr(errc)
	}

	log.Printf("[INFO] Deleting API key: %s", d.Id())

	// Destroy the API key
	err := client.DeleteAPIKey(context.Background(), d.Id())
	if err != nil {
		return diag.Errorf("Error deleting API key: %s", err)
	}

	return nil
}

func resourceMailgunApiKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client, errc := meta.(*Config).GetClient("us")
	if errc != nil {
		return diag.FromErr(errc)
	}

	err := resourceMailgunApiKeyRetrieve(d.Id(), client, d)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMailgunApiKeyRetrieve(id string, client *mailgun.Client, d *schema.ResourceData) error {
	resp, err := client.ListAPIKeys(context.Background(), nil)

	if err != nil {
		return fmt.Errorf("Error retrieving API key list: %s", err)
	}

	var apiKey mtypes.APIKey

	for i, key := range resp {
		if resp[i].ID == id {
			apiKey = key
		}
	}

	if apiKey.ID == "" {
		log.Printf("[DEBUG] API key not found with ID: %s", d.Id())
		return fmt.Errorf("API key not found: %s", id)
	}

	_ = d.Set("requestor", apiKey.Requestor)
	_ = d.Set("secret", apiKey.Secret)

	return nil
}


===== FILE: ./mailgun/resource_mailgun_route.go =====

package mailgun

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mailgun/mailgun-go/v5/mtypes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailgun/mailgun-go/v5"
)

func resourceMailgunRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMailgunRouteCreate,
		Read:          resourceMailgunRouteRead,
		Update:        resourceMailgunRouteUpdate,
		Delete:        resourceMailgunRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceMailgunRouteImport,
		},

		Schema: map[string]*schema.Schema{
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},

			"region": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "us",
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},

			"expression": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"actions": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceMailgunRouteImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	setDefaultRegionForImport(d)

	return []*schema.ResourceData{d}, nil
}

func resourceMailgunRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	opts := mtypes.Route{}

	opts.Priority = d.Get("priority").(int)
	opts.Description = d.Get("description").(string)
	opts.Expression = d.Get("expression").(string)
	actions := d.Get("actions").([]interface{})
	actionArray := []string{}
	for _, i := range actions {
		action := i.(string)
		actionArray = append(actionArray, action)
	}
	opts.Actions = actionArray
	log.Printf("[DEBUG] Route create configuration: %v", opts)

	route, err := client.CreateRoute(context.Background(), opts)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(route.Id)

	log.Printf("[INFO] Route ID: %s", d.Id())

	// Retrieve and update state of route
	_, err = resourceMailgunRouteRetrieve(d.Id(), client, d)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMailgunRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return errc
	}

	opts := mtypes.Route{}

	opts.Priority = d.Get("priority").(int)
	opts.Description = d.Get("description").(string)
	opts.Expression = d.Get("expression").(string)
	actions := d.Get("actions").([]interface{})
	actionArray := []string{}

	for _, i := range actions {
		action := i.(string)
		actionArray = append(actionArray, action)
	}
	opts.Actions = actionArray

	log.Printf("[DEBUG] Route update configuration: %v", opts)

	route, err := client.UpdateRoute(context.Background(), d.Id(), opts)

	if err != nil {
		return err
	}

	d.SetId(route.Id)

	log.Printf("[INFO] Route ID: %s", d.Id())

	// Retrieve and update state of route
	_, err = resourceMailgunRouteRetrieve(d.Id(), client, d)

	if err != nil {
		return err
	}

	return nil
}

func resourceMailgunRouteDelete(d *schema.ResourceData, meta interface{}) error {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return errc
	}

	log.Printf("[INFO] Deleting Route: %s", d.Id())

	// Destroy the route
	err := client.DeleteRoute(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting route: %s", err)
	}

	// Give the destroy a chance to take effect
	return resource.RetryContext(context.Background(), 1*time.Minute, func() *resource.RetryError {
		_, err = client.GetRoute(context.Background(), d.Id())
		if err == nil {
			log.Printf("[INFO] Retrying until route disappears...")
			return resource.RetryableError(
				fmt.Errorf("route seems to still exist; will check again"))
		}
		log.Printf("[INFO] Got error looking for route, seems gone: %s", err)
		return nil
	})
}

func resourceMailgunRouteRead(d *schema.ResourceData, meta interface{}) error {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return errc
	}

	_, err := resourceMailgunRouteRetrieve(d.Id(), client, d)

	if err != nil {
		return err
	}

	return nil
}

func resourceMailgunRouteRetrieve(id string, client *mailgun.Client, d *schema.ResourceData) (*mtypes.Route, error) {

	route, err := client.GetRoute(context.Background(), id)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving route: %s", err)
	}

	_ = d.Set("priority", route.Priority)
	_ = d.Set("description", route.Description)
	_ = d.Set("expression", route.Expression)
	_ = d.Set("actions", route.Actions)

	return &route, nil
}


===== FILE: ./mailgun/resource_mailgun_credential.go =====

package mailgun

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailgun/mailgun-go/v5/mtypes"
)

func resourceMailgunCredential() *schema.Resource {
	log.Printf("[DEBUG] resourceMailgunCredential()")

	return &schema.Resource{
		CreateContext: resourceMailgunCredentialCreate,
		Read:          resourceMailgunCredentialRead,
		Update:        resourceMailgunCredentialUpdate,
		Delete:        resourceMailgunCredentialDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceMailgunCredentialImport,
		},

		Schema: map[string]*schema.Schema{
			"login": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"password": {
				Type:      schema.TypeString,
				ForceNew:  false,
				Required:  true,
				Sensitive: true,
			},

			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"region": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "us",
			},
		},
	}
}

func resourceMailgunCredentialImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	setDefaultRegionForImport(d)

	log.Printf("[DEBUG] Import credential for region '%s' and email '%s'", d.Get("region"), d.Id())

	return []*schema.ResourceData{d}, nil
}

func resourceMailgunCredentialCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	email := fmt.Sprintf("%s@%s", d.Get("login").(string), d.Get("domain").(string))
	password := d.Get("password").(string)

	log.Printf("[DEBUG] Credential create configuration: email: %s", email)

	err := client.CreateCredential(context.Background(), d.Get("domain").(string), email, password)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(email)

	log.Printf("[INFO] Create credential ID: %s", d.Id())

	return nil
}

func resourceMailgunCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return errc
	}

	email := fmt.Sprintf("%s@%s", d.Get("login").(string), d.Get("domain").(string))
	password := d.Get("password").(string)

	log.Printf("[DEBUG] Credential create configuration: email: %s", email)

	err := client.ChangeCredentialPassword(context.Background(), d.Get("domain").(string), email, password)

	if err != nil {
		return err
	}

	d.SetId(email)

	log.Printf("[INFO] Update credential ID: %s", d.Id())

	return nil
}

func resourceMailgunCredentialDelete(d *schema.ResourceData, meta interface{}) error {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return errc
	}

	email := fmt.Sprintf("%s@%s", d.Get("login").(string), d.Get("domain").(string))
	err := client.DeleteCredential(context.Background(), d.Get("domain").(string), email)

	if err != nil {
		return fmt.Errorf("Error deleting credential: %s", err)
	}

	return nil
}

func resourceMailgunCredentialRead(d *schema.ResourceData, meta interface{}) error {
	parts := strings.SplitN(d.Id(), "@", 2)

	if len(parts) != 2 {
		return fmt.Errorf("The ID of credential '%s' don't contains domain!", d.Id())
	}

	login := parts[0]
	domain := parts[1]

	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return errc
	}

	log.Printf("[DEBUG] Read credential for region '%s' and email '%s'", d.Get("region"), d.Id())

	itCredentials := client.ListCredentials(domain, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	var page []mtypes.Credential

	for itCredentials.Next(ctx, &page) {
		log.Printf("[DEBUG] Read credential get new page")

		for _, c := range page {
			if c.Login == d.Id() {
				_ = d.Set("login", login)
				_ = d.Set("domain", domain)
				return nil
			}
		}
	}

	if err := itCredentials.Err(); err != nil {
		return err
	}

	return fmt.Errorf("The credential '%s' not found!", d.Id())
}


===== FILE: ./mailgun/resource_mailgun_domain.go =====

package mailgun

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v5/mtypes"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailgun/mailgun-go/v5"
)

func resourceMailgunDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMailgunDomainCreate,
		ReadContext:   resourceMailgunDomainRead,
		UpdateContext: resourceMailgunDomainUpdate,
		DeleteContext: resourceMailgunDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceMailgunDomainImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"region": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "us",
			},

			"spam_action": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "disabled",
			},

			"smtp_login": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"smtp_password": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  false,
				Sensitive: true,
			},

			"wildcard": {
				Type:     schema.TypeBool,
				ForceNew: true,
				Optional: true,
			},

			"dkim_selector": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"force_dkim_authority": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},

			"open_tracking": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},

			"click_tracking": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},

			"web_scheme": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "http",
			},

			"receiving_records": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "Use `receiving_records_set` instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"receiving_records_set": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      domainRecordsSchemaSetFunc,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"sending_records": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "Use `sending_records_set` instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"sending_records_set": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      domainRecordsSchemaSetFunc,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dkim_key_size": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				if diff.HasChange("name") {
					var sendingRecords []interface{}

					sendingRecords = append(sendingRecords, map[string]interface{}{"id": diff.Get("name").(string)})
					sendingRecords = append(sendingRecords, map[string]interface{}{"id": "_domainkey." + diff.Get("name").(string)})
					sendingRecords = append(sendingRecords, map[string]interface{}{"id": "email." + diff.Get("name").(string)})

					if err := diff.SetNew("sending_records_set", schema.NewSet(domainRecordsSchemaSetFunc, sendingRecords)); err != nil {
						return fmt.Errorf("error setting new sending_records_set diff: %w", err)
					}

					var receivingRecords []interface{}

					receivingRecords = append(receivingRecords, map[string]interface{}{"id": "mxa.mailgun.org"})
					receivingRecords = append(receivingRecords, map[string]interface{}{"id": "mxb.mailgun.org"})

					if err := diff.SetNew("receiving_records_set", schema.NewSet(domainRecordsSchemaSetFunc, receivingRecords)); err != nil {
						return fmt.Errorf("error setting new receiving_records_set diff: %w", err)
					}
				}

				return nil
			},
		),
	}
}

func domainRecordsSchemaSetFunc(v interface{}) int {
	m, ok := v.(map[string]interface{})

	if !ok {
		return 0
	}

	if v, ok := m["id"].(string); ok {
		return stringHashcode(v)
	}

	return 0
}

func resourceMailgunDomainImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	setDefaultRegionForImport(d)

	return []*schema.ResourceData{d}, nil
}

func resourceMailgunDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var name = d.Get("name").(string)
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	var currentData schema.ResourceData
	var newPassword = d.Get("smtp_password").(string)
	var smtpLogin = d.Get("smtp_login").(string)
	var openTracking = d.Get("open_tracking").(bool)
	var clickTracking = d.Get("click_tracking").(bool)
	var webScheme = d.Get("web_scheme").(string)

	// Retrieve and update state of domain
	_, errc = resourceMailgunDomainRetrieve(d.Id(), client, &currentData)

	if errc != nil {
		return diag.FromErr(errc)
	}

	// Update default credential if changed
	if currentData.Get("smtp_password") != newPassword {
		errc = client.ChangeCredentialPassword(ctx, name, smtpLogin, newPassword)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	if currentData.Get("open_tracking") != openTracking {
		var openTrackingValue = "no"
		if openTracking {
			openTrackingValue = "yes"
		}
		errc = client.UpdateOpenTracking(ctx, name, openTrackingValue)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	if currentData.Get("click_tracking") != clickTracking {
		var clickTrackingValue = "no"
		if clickTracking {
			clickTrackingValue = "yes"
		}
		errc = client.UpdateClickTracking(ctx, d.Get("name").(string), clickTrackingValue)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	if currentData.Get("web_scheme") != webScheme {
		opts := mailgun.UpdateDomainOptions{}
		opts.WebScheme = webScheme
		errc = client.UpdateDomain(ctx, name, &opts)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	return nil
}

func resourceMailgunDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	opts := mailgun.CreateDomainOptions{}

	name := d.Get("name").(string)

	opts.SpamAction = mtypes.SpamAction(d.Get("spam_action").(string))
	opts.Password = d.Get("smtp_password").(string)
	opts.Wildcard = d.Get("wildcard").(bool)
	opts.DKIMKeySize = d.Get("dkim_key_size").(int)
	opts.ForceDKIMAuthority = d.Get("force_dkim_authority").(bool)
	opts.WebScheme = d.Get("web_scheme").(string)
	var dkimSelector = d.Get("dkim_selector").(string)
	var openTracking = d.Get("open_tracking").(bool)
	var clickTracking = d.Get("click_tracking").(bool)

	log.Printf("[DEBUG] Domain create configuration: %#v", opts)

	_, err := client.CreateDomain(context.Background(), name, &opts)

	if err != nil {
		return diag.FromErr(err)
	}

	if dkimSelector != "" {
		errc = client.UpdateDomainDkimSelector(ctx, name, dkimSelector)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}
	if openTracking {
		errc = client.UpdateOpenTracking(ctx, name, "yes")

		if errc != nil {
			return diag.FromErr(errc)
		}
	}
	if clickTracking {
		errc = client.UpdateClickTracking(ctx, d.Get("name").(string), "yes")

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	d.SetId(name)

	log.Printf("[INFO] Domain ID: %s", d.Id())

	// Retrieve and update state of domain
	_, err = resourceMailgunDomainRetrieve(d.Id(), client, d)

	if err != nil {
		return diag.FromErr(errc)
	}

	return nil
}

func resourceMailgunDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	log.Printf("[INFO] Deleting Domain: %s", d.Id())

	// Destroy the domain
	err := client.DeleteDomain(context.Background(), d.Id())
	if err != nil {
		return diag.Errorf("Error deleting domain: %s", err)
	}

	// Give the destroy a chance to take effect
	err = resource.RetryContext(ctx, 5*time.Minute, func() *resource.RetryError {
		_, err = client.GetDomain(ctx, d.Id(), nil)
		if err == nil {
			log.Printf("[INFO] Retrying until domain disappears...")
			return resource.RetryableError(
				fmt.Errorf("domain seems to still exist; will check again"))
		}
		log.Printf("[INFO] Got error looking for domain, seems gone: %s", err)
		return nil
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMailgunDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	_, err := resourceMailgunDomainRetrieve(d.Id(), client, d)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMailgunDomainRetrieve(id string, client *mailgun.Client, d *schema.ResourceData) (*mtypes.GetDomainResponse, error) {

	resp, err := client.GetDomain(context.Background(), id, nil)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving domain: %s", err)
	}

	_ = d.Set("name", resp.Domain.Name)
	_ = d.Set("smtp_login", resp.Domain.SMTPLogin)
	_ = d.Set("wildcard", resp.Domain.Wildcard)
	_ = d.Set("spam_action", resp.Domain.SpamAction)
	_ = d.Set("web_scheme", resp.Domain.WebScheme)

	receivingRecords := make([]map[string]interface{}, len(resp.ReceivingDNSRecords))
	for i, r := range resp.ReceivingDNSRecords {
		receivingRecords[i] = make(map[string]interface{})
		receivingRecords[i]["id"] = r.Value
		receivingRecords[i]["priority"] = r.Priority
		receivingRecords[i]["valid"] = r.Valid
		receivingRecords[i]["value"] = r.Value
		receivingRecords[i]["record_type"] = r.RecordType
	}
	_ = d.Set("receiving_records", receivingRecords)
	_ = d.Set("receiving_records_set", receivingRecords)

	sendingRecords := make([]map[string]interface{}, len(resp.SendingDNSRecords))
	for i, r := range resp.SendingDNSRecords {
		sendingRecords[i] = make(map[string]interface{})
		sendingRecords[i]["id"] = r.Name
		sendingRecords[i]["name"] = r.Name
		sendingRecords[i]["valid"] = r.Valid
		sendingRecords[i]["value"] = r.Value
		sendingRecords[i]["record_type"] = r.RecordType

		if strings.Contains(r.Name, "._domainkey.") {
			sendingRecords[i]["id"] = "_domainkey." + resp.Domain.Name
		}
	}
	_ = d.Set("sending_records", sendingRecords)
	_ = d.Set("sending_records_set", sendingRecords)

	info, err := client.GetDomainTracking(context.Background(), id)
	var openTracking = false
	if info.Open.Active {
		openTracking = true
	}
	_ = d.Set("open_tracking", openTracking)

	var clickTracking = false
	if info.Click.Active {
		clickTracking = true
	}
	_ = d.Set("click_tracking", clickTracking)

	return &resp, nil
}


===== FILE: ./mailgun/resource_mailgun_api_key_test.go =====

package mailgun

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mailgun/mailgun-go/v5/mtypes"
)

func TestAccMailgunApiKey_Basic(t *testing.T) {
	var resp mtypes.APIKey
	role := "admin"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunApiKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunApiKeyConfig(role),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMailgunApiKeyExists("mailgun_api_key.foobar", &resp),
					testAccCheckMailgunApiKeyAttributes(role, &resp),
					resource.TestCheckResourceAttr(
						"mailgun_api_key.foobar", "role", role),
					resource.TestCheckResourceAttr(
						"mailgun_api_key.foobar", "description", "Test API key"),
					resource.TestCheckResourceAttr(
						"mailgun_api_key.foobar", "kind", "user"),
				),
			},
		},
	})
}

func testAccCheckMailgunApiKeyDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mailgun_api_key" {
			continue
		}

		client, errc := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])
		if errc != nil {
			return errc
		}

		resp, _ := client.ListAPIKeys(context.Background(), nil)

		for _, key := range resp {
			if key.ID == rs.Primary.ID {
				return fmt.Errorf("API key still exists: %#v", resp)
			}
		}
	}

	return nil
}

func testAccCheckMailgunApiKeyAttributes(id string, APIKey *mtypes.APIKey) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if APIKey.ID != id {
			return fmt.Errorf("Bad ID: %s", APIKey.ID)
		}

		return nil
	}
}

func testAccCheckMailgunApiKeyExists(n string, APIKey *mtypes.APIKey) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No API key ID is set")
		}

		client, errc := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])
		if errc != nil {
			return errc
		}

		resp, err := client.ListAPIKeys(context.Background(), nil)

		if err != nil {
			return err
		}

		for _, key := range resp {
			if key.ID == rs.Primary.ID {
				*APIKey = key
				return nil
			}
		}

		return fmt.Errorf("API key not found")
	}
}

func testAccCheckMailgunApiKeyConfig(id string) string {
	return `
resource "mailgun_api_key" "foobar" {
    id = "` + id + `"
	description = "Test API key"
	role = "admin"
	kind = "user"
}`
}


===== FILE: ./mailgun/resource_mailgun_webhook.go =====

package mailgun

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMailgunWebhook() *schema.Resource {
	log.Printf("[DEBUG] resourceMailgunWebhook()")

	return &schema.Resource{
		CreateContext: resourceMailgunWebhookCreate,
		ReadContext:   resourceMailgunWebhookRead,
		UpdateContext: resourceMailgunWebhookUpdate,
		DeleteContext: resourceMailgunWebhookDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceMailgunWebhookImport,
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "us",
			},

			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"kind": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					allowedKinds := []string{"accepted", "clicked", "complained", "delivered", "opened", "permanent_fail", "temporary_fail", "unsubscribed"}
					matched := false
					for _, kind := range allowedKinds {
						if kind == v {
							matched = true
						}
					}
					if !matched {
						errs = append(errs, fmt.Errorf("kind must be %s", allowedKinds))
					}
					return
				},
			},

			"urls": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceMailgunWebhookImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	setDefaultRegionForImport(d)

	return []*schema.ResourceData{d}, nil
}

func resourceMailgunWebhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	kind := d.Get("kind").(string)
	urls := d.Get("urls").(*schema.Set)

	stringUrls := []string{}
	for _, url := range urls.List() {
		stringUrls = append(stringUrls, url.(string))
	}

	err := client.CreateWebhook(ctx, d.Get("domain").(string), kind, stringUrls)
	if err != nil {
		return diag.FromErr(err)
	}

	id := generateId(d)
	d.SetId(id)

	log.Printf("[INFO] Create webhook ID: %s", d.Id())

	return resourceMailgunWebhookRead(ctx, d, meta)
}

func resourceMailgunWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	kind := d.Get("kind").(string)
	urls := d.Get("urls").(*schema.Set)

	stringUrls := []string{}
	for _, url := range urls.List() {
		stringUrls = append(stringUrls, url.(string))
	}

	err := client.UpdateWebhook(ctx, d.Get("domain").(string), kind, stringUrls)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[INFO] Update webhook ID: %s", d.Id())

	return resourceMailgunWebhookRead(ctx, d, meta)
}

func resourceMailgunWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	kind := d.Get("kind").(string)

	err := client.DeleteWebhook(ctx, d.Get("domain").(string), kind)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[INFO] Delete webhook ID: %s", d.Id())

	return nil
}

func resourceMailgunWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	kind := d.Get("kind").(string)
	urls, err := client.GetWebhook(ctx, d.Get("domain").(string), kind)
	if err != nil {
		return diag.FromErr(err)
	}

	_ = d.Set("kind", kind)
	_ = d.Set("urls", urls)

	return nil
}

func generateId(d *schema.ResourceData) string {
	region := d.Get("region").(string)
	domain := d.Get("domain").(string)
	kind := d.Get("kind").(string)
	return fmt.Sprintf("%s:%s:%s", region, domain, kind)
}


===== FILE: ./mailgun/mailgun_helper.go =====

package mailgun

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"hash/crc32"
	"strings"
)

func setDefaultRegionForImport(d *schema.ResourceData) {
	parts := strings.SplitN(d.Id(), ":", 2)

	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		_ = d.Set("region", "us")
	} else {
		_ = d.Set("region", parts[0])
		d.SetId(parts[1])
	}
}

// stringHashcode hashes a string to a unique hashcode.
//
// crc32 returns an uint32, but for our use we need
// and non-negative integer. Here we cast to an integer
// and invert it if the result is negative.
func stringHashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}


===== FILE: ./mailgun/data_source_mailgun_domain.go =====

package mailgun

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMailgunDomain() *schema.Resource {

	mailgunSchema := resourceMailgunDomain()

	return &schema.Resource{
		Read:   dataSourceMailgunDomainRead,
		Schema: mailgunSchema.Schema,
	}
}

func dataSourceMailgunDomainRead(d *schema.ResourceData, meta interface{}) error {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return errc
	}

	name := d.Get("name").(string)

	_, err := resourceMailgunDomainRetrieve(name, client, d)

	if err != nil {
		return err
	}

	d.SetId(name)
	return nil
}


===== FILE: ./mailgun/resource_mailgun_route_test.go =====

package mailgun

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/mailgun/mailgun-go/v5/mtypes"
)

func TestAccMailgunRoute_Basic(t *testing.T) {
	var route mtypes.Route

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunRouteConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMailgunRouteExists("mailgun_route.foobar", &route),
					resource.TestCheckResourceAttr(
						"mailgun_route.foobar", "priority", "0"),
					resource.TestCheckResourceAttr(
						"mailgun_route.foobar", "description", "inbound"),
					resource.TestCheckResourceAttr(
						"mailgun_route.foobar", "expression", "match_recipient('.*@example.com')"),
					resource.TestCheckResourceAttr(
						"mailgun_route.foobar", "actions.0", "forward('http://example.com/api/v1/foos/')"),
					resource.TestCheckResourceAttr(
						"mailgun_route.foobar", "actions.1", "stop()"),
				),
			},
		},
	})
}

func TestAccMailgunRoute_Import(t *testing.T) {
	resourceName := "mailgun_route.foobar"
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunRouteConfig,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckMailgunRouteDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mailgun_route" {
			continue
		}

		client, _ := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])

		route, err := client.GetRoute(context.Background(), rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Route still exists: %#v", route)
		}
	}

	return nil
}

func testAccCheckMailgunRouteExists(n string, Route *mtypes.Route) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Route ID is set")
		}

		client := testAccProvider.Meta().(*Config)

		err := resource.RetryContext(context.Background(), 1*time.Minute, func() *resource.RetryError {
			var err error
			*Route, err = client.MailgunClient.GetRoute(context.Background(), rs.Primary.ID)

			if err != nil {
				return resource.NonRetryableError(err)
			}

			return nil
		})

		if err != nil {
			return fmt.Errorf("Unable to find Route after retries: %s", err)
		}

		if Route.Id != rs.Primary.ID {
			return fmt.Errorf("Route not found")
		}

		return nil
	}
}

const testAccCheckMailgunRouteConfig = `
resource "mailgun_route" "foobar" {
    priority = "0"
    description = "inbound"
    expression = "match_recipient('.*@example.com')"
    actions = [
        "forward('http://example.com/api/v1/foos/')",
        "stop()"
    ]
}
`


===== FILE: ./mailgun/resource_mailgun_webhook_test.go =====

package mailgun

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-uuid"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccMailgunWebhook_Basic(t *testing.T) {

	uuid, _ := uuid.GenerateUUID()
	domain := fmt.Sprintf("terraformcred.%s.com", uuid)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunWebhookConfig(domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"mailgun_webhook.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"mailgun_webhook.foobar", "region", "us"),
					resource.TestCheckResourceAttr(
						"mailgun_webhook.foobar", "kind", "delivered"),
					resource.TestCheckResourceAttr(
						"mailgun_webhook.foobar", "urls.0", "https://hoge.com"),
				),
			},
		},
	})
}

func TestAccMailgunWebhook_Import(t *testing.T) {
	resourceName := "mailgun_webhook.foobar"
	uuid, _ := uuid.GenerateUUID()
	domain := fmt.Sprintf("terraform.%s.com", uuid)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunWebhookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunWebhookConfig(domain),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckMailgunWebhookDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mailgun_webhook" {
			continue
		}

		client, _ := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])

		kind := rs.Primary.Attributes["kind"]
		webhooks, err := client.GetWebhook(context.Background(), rs.Primary.Attributes["domain"], kind)

		if err == nil {
			return fmt.Errorf("Webhook still exists: %#v", webhooks)
		}
	}

	return nil
}

func testAccCheckMailgunWebhookConfig(domain string) string {
	return `
resource "mailgun_domain" "foobar" {
    name = "` + domain + `"
	spam_action = "disabled"
	region = "us"
    wildcard = true
}

resource "mailgun_webhook" "foobar" {
  domain = mailgun_domain.foobar.id
  region = "us"
  kind = "delivered"
  urls = ["https://hoge.com"]
}`
}


===== FILE: ./mailgun/resource_mailgun_credential_test.go =====

package mailgun

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/mailgun/mailgun-go/v5/mtypes"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccMailgunDomainCredential_Basic(t *testing.T) {

	uuid, _ := uuid.GenerateUUID()
	domain := fmt.Sprintf("terraformcred.%s.com", uuid)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunCrendentialDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunCredentialConfig(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMailgunCredentialExists("mailgun_domain_credential.foobar"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "login", "test_crendential"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "password", "supersecretpassword1234"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "region", "us"),
				),
			},
		},
	})
}

func TestAccMailgunDomainCredential_Update(t *testing.T) {

	uuid, _ := uuid.GenerateUUID()
	domain := fmt.Sprintf("terraform.%s.com", uuid)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunCrendentialDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckMailgunCredentialConfig(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMailgunCredentialExists("mailgun_domain_credential.foobar"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "login", "test_crendential"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "password", "supersecretpassword1234"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "region", "us"),
				),
			},
			{
				Config: testAccCheckMailgunCredentialConfigUpdate(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMailgunCredentialExists("mailgun_domain_credential.foobar"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "domain", domain),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "login", "test_crendential"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "password", "azertyuyiop123456987"),
					resource.TestCheckResourceAttr(
						"mailgun_domain_credential.foobar", "region", "us"),
				),
			},
		},
	})
}

func testAccCheckMailgunCrendentialDestroy(s *terraform.State) error {

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mailgun_domain_credential" {
			continue
		}

		client, err := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])

		resp, err := client.GetDomain(context.Background(), rs.Primary.Attributes["domain"], nil)
		if err == nil {

			itCredentials := client.ListCredentials(rs.Primary.Attributes["domain"], nil)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()

			var page []mtypes.Credential

			for itCredentials.Next(ctx, &page) {

				for _, c := range page {
					if c.Login == rs.Primary.ID {
						return fmt.Errorf("The credential '%s' found! Created at: %s", rs.Primary.ID, c.CreatedAt.String())
					}
				}
			}

			if err := itCredentials.Err(); err != nil {
				return err
			}

			return fmt.Errorf("Domain still exists: %#v", resp)
		}
	}

	return nil
}

func testAccCheckMailgunCredentialExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No domain credential ID is set")
		}

		client, _ := testAccProvider.Meta().(*Config).GetClient(rs.Primary.Attributes["region"])

		itCredentials := client.ListCredentials(rs.Primary.Attributes["domain"], nil)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		var page []mtypes.Credential

		for itCredentials.Next(ctx, &page) {
			for _, c := range page {
				if c.Login == rs.Primary.ID {
					return nil
				}
			}
		}

		if err := itCredentials.Err(); err != nil {
			return err
		}

		return fmt.Errorf("The credential '%s' not found!", rs.Primary.ID)
	}
}

func testAccCheckMailgunCredentialConfig(domain string) string {
	return `
resource "mailgun_domain" "foobar" {
    name = "` + domain + `"
	spam_action = "disabled"
	region = "us"
    wildcard = true
}

resource "mailgun_domain_credential" "foobar" {
	domain = mailgun_domain.foobar.id
	login = "test_crendential"
	password = "supersecretpassword1234"
	region = "us"
}`
}

func testAccCheckMailgunCredentialConfigUpdate(domain string) string {
	return `
resource "mailgun_domain" "foobar" {
    name = "` + domain + `"
	spam_action = "disabled"
	region = "us"
    wildcard = true
}

resource "mailgun_domain_credential" "foobar" {
	domain = mailgun_domain.foobar.id
	login = "test_crendential"
	password = "azertyuyiop123456987"
	region = "us"
}`
}


===== FILE: ./mailgun/data_source_mailgun_domain_test.go =====

package mailgun

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccMailgunDomainDataSource_Basic(t *testing.T) {
	uuid, _ := uuid.GenerateUUID()
	domain := fmt.Sprintf("terraform.%s.com", uuid)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: newProvider(),
		CheckDestroy:      testAccCheckMailgunDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMailgunDomainDataSourceConfig_Basic(domain),
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceMailgunDomainCheck("data.mailgun_domain.test", domain),
				),
			},
		},
	})
}

func testAccDataSourceMailgunDomainCheck(name string, domain string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s\n%#v", name, rs)
		}

		attr := rs.Primary.Attributes

		if attr["name"] != domain {
			return fmt.Errorf("bad name %s", attr["name"])
		}

		if attr["spam_action"] != "disabled" {
			return fmt.Errorf("Bad spam_action: %s", attr["spam_action"])
		}

		if attr["wildcard"] != "false" {
			return fmt.Errorf("Bad wildcard: %s", attr["wildcard"])
		}

		return nil
	}
}

func testAccMailgunDomainDataSourceConfig_Basic(domain string) string {
	return fmt.Sprintf(`
resource "mailgun_domain" "foobar" {
	name = "%s"
	spam_action = "disabled"
	wildcard = false
}
data "mailgun_domain" "test" {
	name = mailgun_domain.foobar.id
}
`, domain)
}


===== FILE: ./.travis.yml =====

dist: trusty
sudo: required
services:
- docker
language: go
go:
- 1.24.1

install:
# This script is used by the Travis build to install a cookie for
# go.googlesource.com so rate limits are higher when using `go get` to fetch
# packages that live there.
# See: https://github.com/golang/go/issues/12933
- bash scripts/gogetcookie.sh
- go get github.com/kardianos/govendor

script:
- make test
- make vendor-status
- make vet
- make website-test

branches:
  only:
  - master
matrix:
  fast_finish: true
  allow_failures:
  - go: tip


===== FILE: ./main.go =====

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-providers/terraform-provider-mailgun/mailgun"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mailgun.Provider})
}
