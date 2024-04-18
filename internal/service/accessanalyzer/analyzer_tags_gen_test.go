// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package accessanalyzer_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func testAccAccessAnalyzerAnalyzer_tagsSerial(t *testing.T) {
	t.Helper()

	t.Run("basic", testAccAccessAnalyzerAnalyzer_tags)
	t.Run("null", testAccAccessAnalyzerAnalyzer_tags_null)
	t.Run("AddOnUpdate", testAccAccessAnalyzerAnalyzer_tags_AddOnUpdate)
	t.Run("EmptyTag_OnCreate", testAccAccessAnalyzerAnalyzer_tags_EmptyTag_OnCreate)
	t.Run("EmptyTag_OnUpdate_Add", testAccAccessAnalyzerAnalyzer_tags_EmptyTag_OnUpdate_Add)
	t.Run("EmptyTag_OnUpdate_Replace", testAccAccessAnalyzerAnalyzer_tags_EmptyTag_OnUpdate_Replace)
	t.Run("DefaultTags_providerOnly", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_providerOnly)
	t.Run("DefaultTags_nonOverlapping", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_nonOverlapping)
	t.Run("DefaultTags_overlapping", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_overlapping)
	t.Run("DefaultTags_updateToProviderOnly", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_updateToProviderOnly)
	t.Run("DefaultTags_updateToResourceOnly", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_updateToResourceOnly)
	t.Run("DefaultTags_emptyResourceTag", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_emptyResourceTag)
	t.Run("DefaultTags_nullOverlappingResourceTag", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_nullOverlappingResourceTag)
	t.Run("DefaultTags_nullNonOverlappingResourceTag", testAccAccessAnalyzerAnalyzer_tags_DefaultTags_nullNonOverlappingResourceTag)
}

func testAccAccessAnalyzerAnalyzer_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAnalyzerConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAnalyzerConfig_tags0(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_null(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyzerConfig_tagsNull(rName, "key1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config:             testAccAnalyzerConfig_tags0(rName),
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_AddOnUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyzerConfig_tags0(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_EmptyTag_OnCreate(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAnalyzerConfig_tags0(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_EmptyTag_OnUpdate_Add(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				Config: testAccAnalyzerConfig_tags2(rName, "key1", "value1", "key2", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_EmptyTag_OnUpdate_Replace(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", ""),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_providerOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccAnalyzerConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags2("key1", "value1updated", "key2", "value2"),
					testAccAnalyzerConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key2", "value2"),
					testAccAnalyzerConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key2", "value2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags0(),
					testAccAnalyzerConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
					testAccAnalyzerConfig_tags1(rName, "resourcekey1", "resourcevalue1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1updated"),
					testAccAnalyzerConfig_tags2(rName, "resourcekey1", "resourcevalue1updated", "resourcekey2", "resourcevalue2"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.resourcekey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey1", "resourcevalue1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.resourcekey2", "resourcevalue2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags0(),
					testAccAnalyzerConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_overlapping(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
					testAccAnalyzerConfig_tags1(rName, "overlapkey1", "resourcevalue1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags2("overlapkey1", "providervalue1", "overlapkey2", "providervalue2"),
					testAccAnalyzerConfig_tags2(rName, "overlapkey1", "resourcevalue1", "overlapkey2", "resourcevalue2"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey2", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey2", "resourcevalue2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("overlapkey1", "providervalue1"),
					testAccAnalyzerConfig_tags1(rName, "overlapkey1", "resourcevalue2"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.overlapkey1", "resourcevalue2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.overlapkey1", "resourcevalue2"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_updateToProviderOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccAnalyzerConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_updateToResourceOnly(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccAnalyzerConfig_tags0(rName),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				Config: testAccAnalyzerConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_emptyResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "value1"),
					testAccAnalyzerConfig_tags1(rName, "key1", ""),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", ""),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_nullOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("key1", "providervalue1"),
					testAccAnalyzerConfig_tagsNull(rName, "key1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.key1", "providervalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccAccessAnalyzerAnalyzer_tags_DefaultTags_nullNonOverlappingResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.AnalyzerSummary
	resourceName := "aws_accessanalyzer_analyzer.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.AccessAnalyzerServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckAnalyzerDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: acctest.ConfigCompose(
					acctest.ConfigDefaultTags_Tags1("providerkey1", "providervalue1"),
					testAccAnalyzerConfig_tagsNull(rName, "resourcekey1"),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckAnalyzerExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.providerkey1", "providervalue1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}